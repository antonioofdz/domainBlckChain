
# .PHONY: protos

contract:
	docker-compose run solc --abi ./contracts/tienda.sol -o ./contracts --overwrite
