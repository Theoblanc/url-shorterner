test:
	go test ./...


redis:
	docker run --name redis -d -p 6379:6379 redis:alpine


docker:
	docker run --name url-shortener -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres 