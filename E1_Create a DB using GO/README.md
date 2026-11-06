# Create a PostgreSQL Database with Go

A small Go program that connects to a PostgreSQL server and creates a database if it does not already exist.

## How it works

1. Validates the database name (letters, digits and underscores only) to prevent SQL injection.
2. Connects to the default `postgres` maintenance database. The connection URL is built with `net/url`, so special characters in the password are escaped automatically.
3. Checks whether the database exists by querying the `pg_database` system catalog.
4. Runs `CREATE DATABASE` if it does not exist.

## Requirements

- [Go](https://go.dev/dl) 1.21 or newer
- [PostgreSQL](https://www.postgresql.org/download/) running locally (or reachable over the network)

## Project files

All files are inside the `E1_Create a DB using GO` folder.

|File|Purpose|
|---|---|
|`main.go`|The program|
|`go.mod`|Module name and dependency versions|
|`go.sum`|Dependency checksums|

## Configuration

Edit these variables at the top of `main()` in `main.go`:

|Variable|Description|Default|
|---|---|---|
|`host`|Server hostname|`localhost`|
|`port`|Server port|`5432`|
|`user`|PostgreSQL user|`postgres`|
|`password`|Password for that user|`YOUR_PASSWORD`|
|`databaseName`|Database to create|`new_database`|

## Usage

This program lives in the `E1_Create a DB using GO` folder of the [DB_Lab](https://github.com/Soheil-Latifi-04/DB_Lab) repository. Clone the repository, move into the folder and download the dependencies:

```
git clone https://github.com/Soheil-Latifi-04/DB_Lab.git
cd DB_Lab
cd "E1_Create a DB using GO"
go mod tidy
```

Run the program:

```
go run main.go
```

Expected output on the first run:

```
Database new_database created successfully.
```

Running it again prints:

```
Database new_database already exists.
```

### Starting from `main.go` only

If the repository does not contain `go.mod` and `go.sum`, generate them first:

```
go mod init pgdemo
go get github.com/jackc/pgx/v5
```

## Verify the result

```
psql -U postgres -h localhost -p 5432 -l
```

The new database should appear in the list. You can also check in pgAdmin under **Databases**.

## Troubleshooting

|Error|Likely cause and fix|
|---|---|
|`connection refused`|The PostgreSQL service is not running or the port is wrong. Start the service and check `port` in `main.go`.|
|`password authentication failed`|Wrong password. Test it with `psql` to confirm.|
|`invalid database name`|The name contains characters other than letters, digits and underscores, or starts with a digit.|
|`go get` fails or times out|The Go package proxy may be blocked. Try `go env -w GOPROXY=https://goproxy.io,direct`.|

## Notes

- PostgreSQL has no `CREATE DATABASE IF NOT EXISTS`, so the program checks for the database before creating it.
- Database names are inserted into the `CREATE DATABASE` statement directly because SQL placeholders only work for values, not identifiers. The name validation at the start of the program makes this safe.
- Names are double-quoted, so PostgreSQL keeps their exact case. Lowercase names are the easiest to work with.