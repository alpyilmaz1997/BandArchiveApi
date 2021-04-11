.PHONY:check_install
check_install:
	which swagger || GO111MODULE=off go get -u github.com/goswagger/goswagger

.PHONY:swagger
swagger: check_install
	GO111MODULE=off swagger generate spec -o /swagger.yaml --scan-models