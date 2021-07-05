# Challenges

- ~~Create an "updateArticle" route, parse HTTP request body + find Article in array and update the entry~~
- Connect to MySQL DB
- Prevent race conditions using Mutex: https://tutorialedge.net/golang/go-mutex-tutorial/
- Add tests
- Add authorization?
- Refactor Article to separate file / module?

## Migrations
Replace with DB driver of choice
Replace migration directory with own absolute local URL

```docker run -v /Users/ericlandheer/Programming/Go/rest-api/migrations:/migrations --network host migrate/migrate -path=./migrations/ -database "mysql://go_mysql:test@tcp(localhost:3306)/go" up```