VERSION = $(shell cat .version)
GHT = $(GITHUB_TOKEN)

release: pf2s3
	ghr  --username baldguysoftware --token ${GITHUB_TOKEN} --replace ${VERSION} pf2s3


pf2s3: 
	go get -t ./...
	go build 

test:
	go test
	go vet

