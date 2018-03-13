VERSION = $(shell cat .version)
GHT = $(GITHUB_TOKEN)

release: pf2s3
	ghr  --username therealbill --token ${GITHUB_TOKEN} --replace ${VERSION} pf2s3


pf2s3: 
	go build 

test:
	go test
	go vet

