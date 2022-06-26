run:
	go run ./src/main.go
docker-build:
	docker build -t github.com/tomkaith13/fhir-questionnaire-engine .

docker-run:
	docker run -p 8080:8080 github.com/tomkaith13/fhir-questionnaire-engine

docker-image-clean:
	docker stop github.com/tomkaith13/fhir-questionnaire-engine && docker rm -f github.com/tomkaith13/fhir-questionnaire-engine