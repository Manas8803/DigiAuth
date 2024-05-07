.PHONY: build deploy clean all all-swap

run:
	cd app && npm run go

install:
	cd app && npm i	

build:
	GOOS=linux GOARCH=amd64 go build -o ./app/bootstrap ./app/cmd/main.go  

deploy:
	cd deploy-scripts && cdk deploy

deploy-swap:
	cd deploy-scripts && cdk deploy --hotswap

clean:
	rm -rf ./app/bootstrap
	
all:
	make clean
	make build
	make deploy

all-swap:
	make clean
	make build
	make deploy-swap