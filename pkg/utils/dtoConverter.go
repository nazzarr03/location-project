package utils

import "encoding/json"

func DTOtoJSON(dto interface{}, data interface{}) error {
	b, err := json.Marshal(&dto)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return nil
}

func JSONtoDTO(data interface{}, dto interface{}) error {
	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &dto)
	if err != nil {
		return err
	}
	return nil
}
