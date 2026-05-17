<!-- distilled: 2026-05-12T09:16:46.943Z | source: hello.go -->

<!-- distilled: 2026-05-12T09:15:59.898Z | source: hello.go -->

<!-- distilled: 2026-05-12T09:09:57.340Z | source: hello.go -->

This file implements a simple multilingual "Hello" greeting function in the `main` package. It defines unexported constants for language names (`spanish`, `french`, `german`) and their corresponding greeting prefixes. The key function `Hello(name, language string) string` returns a personalized greeting by combining the appropriate language prefix with the provided name (defaulting to "World" if empty). It delegates prefix lookup to `greetingPrefix(language string) (prefix string)`, which uses a switch statement. The file imports `fmt` and runs via `main()`, printing `"Hello, world"` to stdout.