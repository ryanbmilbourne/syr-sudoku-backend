
CONTAINER := syr-sudoku-backend

.PHONY: container
container: 
	docker build -t $(CONTAINER) -f Dockerfile .

.PHONY: runlocal
runlocal: container
	docker run -it --env-file .env -p 8080:8080 $(CONTAINER)


.PHONY: deploydev
deploydev:
	git checkout master
	git pull
	heroku container:push web --app sudoku-dev
