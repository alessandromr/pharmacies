GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')
TEST?=./...

fmt:
	gofmt -w $(GOFMT_FILES)
	go vet ./...

integration:
	docker-compose down
	docker-compose up -d
	sleep 2
	LOCAL_TESTING=true go test $(TEST) -run Integration
	# docker-compose down

dao_repeat:
	echo "This test require some seconds (4-10s)"
	docker-compose down
	docker-compose up -d
	sleep 2
	LOCAL_TESTING=true go test ./dao/ -run Integration -count 100
	docker-compose down
	