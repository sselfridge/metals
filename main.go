package main

import (
	"reflect"
	"fmt"
	// "encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)



func add(a int64, b int64) int64 {
	if b == 0 {
		return a
	}

	return add(a^b, (a&b)<<1)
}

//for mupltiplcation and division
//make both numbers positive and return a flag if their product needs to be inverted
func negCheck(a int64, b int64) (int64, int64, bool) {
	makeNeg := false
	if a < 0 {
		makeNeg = true
		a = invert(a)
	}
	if b < 0 && makeNeg {
		makeNeg = false
		b = invert(b)
	} else if b < 0 && !makeNeg {
		b = invert(b)
		makeNeg = true
	}
	return a, b, makeNeg
}

func multi(a int64, b int64) int64 {

	a, b, makeNeg := negCheck(a, b)

	sum := int64(0)
	for i := int64(1); i <= b; i = add(i, int64(1)) {
		sum = add(sum, a)
	}

	if makeNeg {
		sum = invert(sum)
	}
	return sum
}

func invert(num int64) int64 {
	return add(^num, int64(1))
}

func sub(a int64, b int64) int64 {
	b = invert(b)
	return add(a, b)
}

func div(a int64, b int64) (int64, int64) {
	a, b, makeNeg := negCheck(a, b)

	divCount := int64(0)
	remainder := int64(0)
	for a > 0 {
		a = sub(a, b)
		divCount = add(divCount, int64(1))
	}
	fmt.Println(a)
	if a != 0 {
		divCount = sub(divCount, int64(1))
		remainder = invert(a)
	}

	if makeNeg {
		divCount = invert(divCount)
	}

	return divCount, remainder

}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO check that query strings exist and their type, throw errors otherwise

	num1, err := strconv.ParseInt(request.QueryStringParameters["num1"], 10, 64)
	fmt.Println(num1, err, reflect.TypeOf(num1))

	num2, err := strconv.ParseInt(request.QueryStringParameters["num2"], 10, 64)
	fmt.Println(num2, err, reflect.TypeOf(num2))

	op := request.QueryStringParameters["op"]

	answer := int64(0)
	remainder := int64(0)
	switch string := op; string {
	case "add":
		answer = add(num1, num2)
	case "multi":
		answer = multi(num1, num2)
	case "sub":
		answer = sub(num1, num2)
	case "div":
		answer, remainder = div(num1, num2)
		//default //add error
	}

	output := fmt.Sprintln(answer, "r", remainder)
	resp := events.APIGatewayProxyResponse{StatusCode: 200, Body: output}
	resp.Headers = make(map[string]string)

	// These headers show up in AWS test console and in POSTMAN but
	// still can't make the GET request from the browser as needed
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Credentials"] = "true"
	resp.Headers["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	resp.Headers["Access-Control-Allow-Methods"] = "GET"

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
