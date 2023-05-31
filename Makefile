run:
	go run ./src/main.go
dbuild:
	docker build -t github.com/tomkaith13/fhir-questionnaire-engine .

drun:
	make dbuild && docker run -p 8080:8080 github.com/tomkaith13/fhir-questionnaire-engine

clean:
	docker stop github.com/tomkaith13/fhir-questionnaire-engine && docker rm -f github.com/tomkaith13/fhir-questionnaire-engine