# Go notes
Sample "Notes" application for golang learning purposes.

Implemented different notes storage engines:
* memory (when server if off notes are lost)
* file (json format stored in local file, use config.yaml to specify file)

Run all test:
 **`go test ./...`**

Run web server `go run main.go`