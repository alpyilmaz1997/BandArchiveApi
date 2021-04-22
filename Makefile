.PHONY:swagger
swagger: 
	swagger generate server -f ./swagger.yaml -A bandArchiceApi
