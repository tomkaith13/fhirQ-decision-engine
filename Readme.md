# FHIR Questionnaire Decision Engine

## Summary
This dockerized web-app will be used to deploy
code that will take in a series of FHIR Questionnaire
and once processed, consumer of the app can 
call into the app by sending a FHIR QuestionnaireResponse
and get next decision node

The decision engine template is completely dockerized.

## How to Run via docker
- Run `make docker-build` to build the container
- Run `make docker-run` to run the container which will expose the 8080 port


