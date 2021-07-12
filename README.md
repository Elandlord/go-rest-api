# Challenges

- ~~Create an "updateArticle" route, parse HTTP request body + find Article in array and update the entry~~
- ~~Connect to MySQL DB~~
- ~~Refactor Article to separate file / module?~~
- ~~Prevent race conditions using Mutex: https://tutorialedge.net/golang/go-mutex-tutorial/~~ **Not necessary, the RDBMS should handle this.** 
- ~~Add authorization - _Active, using JWT_~~
- Add tests
- Replace hard coded login credentials with table credentials
- Deploy using Docker (Kubernetes)


## Steps

- Make sure you have Docker installed, and run `docker-compose up`.

## Migrations
Replace with DB driver of choice
Replace migration directory with own absolute local URL

```docker run -v /Users/ericlandheer/Programming/Go/rest-api/migrations:/migrations --network host migrate/migrate -path=./migrations/ -database "mysql://go_mysql:test@tcp(localhost:3306)/go" up```
