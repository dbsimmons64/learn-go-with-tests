# Why the second test doesn't time out after 20ms

## Observation
The test takes ~12 seconds to complete instead of ~20ms. The log shows:
```
httptest.Server blocked in Close after 5 seconds, waiting for connections:
  *net.TCPConn ... in state active
```

## Root Cause

The `ping()` function spawns a goroutine that calls `http.Get(url)` with **no timeout**. When the select's `time.After(20ms)` case fires, `ConfigurableRacer` returns the error immediately. But:

1. The goroutine from `ping(serverA)` is still blocked on `http.Get()` — it will wait the full 11 seconds
2. The goroutine from `ping(serverB)` is still blocked on `http.Get()` — it will wait the full 12 seconds
3. These goroutines hold active TCP connections to the test servers
4. When `defer serverA.Close()` and `defer serverB.Close()` run, they wait for all active connections to drain
5. This blocks the test for ~12 seconds

## The Bug in `ping()`

`http.Get(url)` is a blocking call with no deadline. Even though the select statement returns early, the goroutine continues running until the HTTP request completes (or fails). There is no mechanism to cancel or abort the in-flight request when the timeout fires.

## The Fix (not implementing — just explaining)

The `ping()` function needs to use an HTTP client with a timeout/deadline, or the request needs to be cancellable (e.g., using `http.NewRequest` with a context that gets cancelled when the timeout fires).
