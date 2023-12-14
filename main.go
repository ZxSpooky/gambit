package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ZxSpooky/gambit/awsgo"

	"os"
	"github.com/ZxSpooky/gambit/bd"
	"strings"
	/*

	"github.com/ZxSpooky/gambit/awsgo"
	
	"github.com/ZxSpooky/gambit/handlers" */

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros(){
		panic("Error en los parametros. debe enviar 'SecretName', 'UserPoolID', 'Region', 'UrlPrefix'")
	}
	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method:= request.RequestContext.HTTP.Method
	body:= request.Body
	header:= request.Headers

	bd.ReadSecret()

	headersResp := map[string]string{
	"Content-Type":"application/json",
	}
	res =&events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}
	return res, nil
}

func ValidoParametros () bool{
	
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro{
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UserPoolId")
	if !traeParametro{
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("Region")
	if !traeParametro{
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro{
		return traeParametro
	}
	
	return traeParametro
}
