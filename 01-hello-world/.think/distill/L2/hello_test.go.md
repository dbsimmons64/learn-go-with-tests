<!-- distilled: 2026-05-12T09:15:59.898Z | source: hello_test.go -->

<!-- distilled: 2026-05-12T09:10:00.817Z | source: hello_test.go -->

- **Purpose**: Tests the `Hello` function for generating localized greetings across multiple languages and edge cases.
- **Key exports**: `TestHello` (table-driven test with subtests for Dave, empty string, Spanish, French, and German) and `assertCorrectMessage` (a test helper that compares got vs. want strings).
- **Imports**: `testing` package for Go's standard testing framework.
- **Usage**: Imported by the `main` package; run via `go test` to validate greeting output.