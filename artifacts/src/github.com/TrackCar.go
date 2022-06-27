package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
	State  string `json:"state"`
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko", State: "Created"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad", State: "Created"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo", State: "Created"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max", State: "Created"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana", State: "Created"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel", State: "Created"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav", State: "Created"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari", State: "Created"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria", State: "Created"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro", State: "Created"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
		
	}

	return nil
}

func (s *SmartContract) ManufactureCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string, state string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
		State: state,
	}
	car.State = "Created"
	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

func (s *SmartContract) QueryCarState(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

func (s *SmartContract) ManufactureToDealer(ctx contractapi.TransactionContextInterface, carNumber string) error {
	car, err := s.QueryCarState(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.State == "Ready_FOR_SALE" {
		return errors.New(" ERROR: CAR IS ALREADY AVAILABLE AT DEALER!")
	}
	if car.State == "Created" {
	
		car.Owner = "Dealer"
		car.State = "Ready_FOR_SALE"
		carAsBytes, _ := json.Marshal(car)

		return ctx.GetStub().PutState(carNumber, carAsBytes)

	} else if car.State == "SOLD" {
			return errors.New("ERROR: THIS CAR NUMBER ID IS ALREADY SOLD TO CUSTOMER BY DEALER!")
		} else {	
		fmt.Printf("CAR IS NOT MANUFACTURED YET!")
		return errors.New("ERROR: CAR IS NOT MANUFACTURED YET!")
	}
	
}

func (s *SmartContract) DealerToCustomer(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	car, err := s.QueryCarState(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.State == "SOLD" {
		return errors.New(" ERROR: CAR IS ALREADY SOLD BY DEALER TO CUSTOMER, PLEASE ASK FOR ANOTHER COPY OF CAR!")
	}

	if car.State == "Ready_FOR_SALE" { 

		car.Owner = newOwner
		car.State = "SOLD"
	    carAsBytes, _ := json.Marshal(car)

	    return ctx.GetStub().PutState(carNumber, carAsBytes)
	} else {
		
		fmt.Printf("CAR IS NOT AVAILABVLE AT DELIVER")
		return errors.New(" ERROR: CAR IS NOT AVAILABVLE AT DELIVER")
	}

	
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create Trackcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting Trackcar chaincode: %s", err.Error())
	}
}







// func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
// 	startKey := "CAR0"
// 	endKey := "CAR99"

// 	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	results := []QueryResult{}

// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()

// 		if err != nil {
// 			return nil, err
// 		}

// 		car := new(Car)
// 		_ = json.Unmarshal(queryResponse.Value, car)

// 		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
// 		results = append(results, queryResult)
// 	}

// 	return results, nil
// }