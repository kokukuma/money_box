PORT = 5000

all: build

.PHONY: build
build:  ## make contract and abi
	solc --abi --bin sol/*.sol -o bin --overwrite
	abigen --abi="bin/MoneyBox.abi" --bin="bin/MoneyBox.bin" --pkg=money_box --out="contract/money_box.go"
	cp bin/MoneyBox.abi frontend/src/assets/MoneyBox.json

.PHONY: run
run:  ## run contract
	go run cmd/run/run.go

.PHONY: deploy
deploy: ## deploy contract
	go run cmd/deploy/deploy.go

.PHONY: serve
serve: ## deploy contract
	cd frontend; npm run serve
