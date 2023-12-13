package converter

import (
	"encoding/json"
	"fmt"
)

// InterfaceToSlice interface类型转slice
func InterfaceToSlice[T any](i interface{}) ([]T, error) {
	s := make([]T, 0)
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert iterface to slice error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &s); err != nil {
		return nil, fmt.Errorf("convert iterface to slice error. %s", err.Error())
	}
	return s, nil
}

// InterfaceToStruct interface  类型转 struct
func InterfaceToStruct[T any](i interface{}) (*T, error) {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert interface to struct error. %s", err.Error())
	}

	t := new(T)
	if err := json.Unmarshal(marshal, t); err != nil {
		return nil, fmt.Errorf("convert interface to struct error. %s", err.Error())
	}
	return t, nil
}

// InterfaceToMap interface 转 map
func InterfaceToMap[K comparable, V any](i interface{}) (map[K]V, error) {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert interface to map error. %s", err.Error())
	}
	m := make(map[K]V, 0)
	if err := json.Unmarshal(marshal, &m); err != nil {
		return nil, fmt.Errorf("convert interface to map error. %s", err.Error())
	}
	return m, nil
}

// MapToStruct map to struct
func MapToStruct[T any](m map[string]interface{}) (*T, error) {
	t := new(T)
	marshal, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("convert map to struct error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &t); err != nil {
		return nil, fmt.Errorf("convert map to struct error. %s", err.Error())
	}
	return t, nil

}

// StructToMap struct to map
func StructToMap[K comparable, V any](s interface{}) (map[K]V, error) {
	m := make(map[K]V, 0)
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("convert struct to map error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &m); err != nil {
		return nil, fmt.Errorf("convert struct to map error. %s", err.Error())
	}
	return m, nil
}
