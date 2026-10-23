# qa-function-go

Example repository for publishing an HTTP function: Go, with Google's Functions
Framework, built with Google's buildpacks.

- The function to run is `HelloHTTP`, the name it is registered with in `function.go` (`functions.HTTP`): write it in Configure.
- It answers every request with a fixed text, so a 200 shows it is up.
