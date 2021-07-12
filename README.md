# Challenges

- ~~Create an "updateArticle" route, parse HTTP request body + find Article in array and update the entry~~
- ~~Connect to MySQL DB~~
- ~~Refactor Article to separate file / module?~~
- ~~Prevent race conditions using Mutex: https://tutorialedge.net/golang/go-mutex-tutorial/~~ **Not necessary, the RDBMS should handle this.** 
- ~~Add JWT authorization~~
- ~~Add tests - _Active_~~
- Replace hard coded login credentials with table credentials
- Deploy using Docker (Kubernetes)


## Steps

- Checkout project `git clone https://github.com/Elandlord/go-rest-api.git`
- Run `cp .env.example .env` and set correct credentials in the `.env`
- Make sure you have Docker installed, and run `docker-compose up`.

## Testing
- Run `go test` to run the tests (only in root directory)
- Run `go test ./...` to run all tests (including subdirectories)

## Migrations
Replace with DB driver of choice
Replace migration directory with own absolute local URL

```docker run -v /Users/ericlandheer/Programming/Go/rest-api/migrations:/migrations --network host migrate/migrate -path=./migrations/ -database "mysql://go_mysql:test@tcp(localhost:3306)/go" up```
