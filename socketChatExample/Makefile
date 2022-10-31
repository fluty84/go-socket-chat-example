.PHONY: up
up:
	go mod vendor
	docker-compose -f docker-compose.yml up -d --build --always-recreate-deps
.PHONY: upnodaemon
upnodaemon:
	go mod vendor
	docker-compose -f docker-compose.yml up --build --always-recreate-deps
.PHONY: down
down:
	docker-compose down