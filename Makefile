.PHONY:swagger
swagger: 
	rm -rf models
	swagger generate server -f ./swagger.yaml -A bandArchiceApi