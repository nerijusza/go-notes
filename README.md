# Go notes
Sample "Notes" application for golang learning purposes.

Implemented different notes storage engines:
* memory (when server if off notes are lost)
* file (json format stored in local file, use config.yaml to specify file)
* airtable.com API need create table with structure (ID int, Content string, DateCreated string) and update config.yaml with correct values (left mine in repo)

Configuration in file: `config.yaml`

Run all test:
 `go test ./...`

Run web server `go run main.go`