package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/ZxSpooky/gambit/awsgo"
	"github.com/ZxSpooky/gambit/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// funcion SecretManager
func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println("> Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	//codifica a json
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura Secret OK" + nombreSecret)
	return datosSecret, nil
}
