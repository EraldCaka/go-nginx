product:
	@go build -o services/product products_service/main.go

user:
	@go build -o services/user users_service/main.go

runp: product
	@ ./services/product

runu: user
	@ ./services/user -port=8083

runu1: user
	@ ./services/user -port=8081
