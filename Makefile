build:
	docker build -t parsyl_challenge_db -f ./database/database.Dockerfile .

bootstrap: build
	make start

start:
	docker run -d -p 5432:5432 --name db parsyl_challenge_db

stop:
	docker stop db
