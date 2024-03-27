package db

import (
    "fmt"
    "errors"
    "log"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/dynamodb"


)



func createDynamoDBSession() (*dynamodb.DynamoDB) {  


    os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
    os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
    os.Setenv("AWS_SESSION_TOKEN", "dummy3")
    cred := credentials.NewEnvCredentials()
    cred.Get()

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
        Endpoint: aws.String("http://localhost:8000"),
        Credentials: cred,
    })
    if err != nil {
        log.Fatalf("Error: %s\n", err)
    }

    return dynamodb.New(sess)
}





func ListTables() (error)  {
    dynamoDBSess := createDynamoDBSession()
    input := &dynamodb.ListTablesInput{}

    fmt.Printf("Tables:\n")

    for {
        // Get the list of tables
        result, err := dynamoDBSess.ListTables(input)
        if err != nil {
            if aerr, ok := err.(awserr.Error); ok {
                switch aerr.Code() {
                case dynamodb.ErrCodeInternalServerError:
                    return errors.New(fmt.Sprintln(dynamodb.ErrCodeInternalServerError, aerr.Error()))
                default:
                    return errors.New(fmt.Sprintln(aerr.Error()))
            }
            } else {
                // Print the error, cast err to awserr.Error to get the Code and
                // Message from an error.
                return errors.New(fmt.Sprintln(err.Error()))
            }
        }

        for _, n := range result.TableNames {
            fmt.Println(*n)
        }

        // assign the last read tablename as the start for our next call to the ListTables function
        // the maximum number of table names returned in a call is 100 (default), which requires us to make
        // multiple calls to the ListTables function to retrieve all table names
        input.ExclusiveStartTableName = result.LastEvaluatedTableName

        if result.LastEvaluatedTableName == nil {
            break
        }
    }

    return nil
}

