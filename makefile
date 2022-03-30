
# .PHONY: protos

contract:
	docker-compose run solc --abi ./contracts/HelloWorld.sol -o ./contracts --optimize --overwrite
	docker-compose run eth abigen --abi ./contracts/HelloWorld.abi --pkg contracts --type Contract --out ./backend/internal/contracts/HelloWorld.go
