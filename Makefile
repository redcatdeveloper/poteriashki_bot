build:
	cd src && go env -w GO111MODULE=on
	cd src && go mod tidy
	cd src && go mod vendor
	cd src && go mod download
	cd src && go build -o ../poteriashki_bot
