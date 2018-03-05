package concataddress

import "fmt"

// Process ...
func Process(input Input) Output {
	result := []concataddress{}
	for _, payload := range input.Payload {
		if payload.Workflow == "completed" && payload.Kind == "htv" {
			result = append(result, concataddress{
				fmt.Sprintf("%s %s %s %s %s", payload.Address.BuildingNumber, payload.Address.Street, payload.Address.Suburb, payload.Address.State, payload.Address.Postcode),
				payload.Kind,
				payload.Workflow,
			})
		}
	}

	return Output{
		Response: result,
	}
}
