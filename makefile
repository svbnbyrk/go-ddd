oapi:
	oapi-codegen -generate types,chi-server -o internal/ports/api.gen.go -package ports api/wallet.yml