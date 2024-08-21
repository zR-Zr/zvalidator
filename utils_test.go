package zvalidator

import (
	"reflect"
	"testing"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type User struct {
	Name            string   `json:"name"`
	Age             *int     `json:"age"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	ConfirmPassword string   `json:"confirm_password"`
	City            string   `json:"city"`
	Address         Address  `json:"address"`
	PtrAddress      *Address `json:"ptr_address"`
}

func TestStructToMap(t *testing.T) {
	age := 25
	ptrAddress := &Address{City: "上海", Street: "花山路"}
	u := User{
		Name:            "John",
		Age:             &age,
		Email:           "john.doe@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		City:            "北京",
		Address:         Address{City: "Beijing", Street: "Long St"},
		PtrAddress:      ptrAddress,
	}

	expectedMap := map[string]any{
		"name":             "John",
		"age":              25,
		"email":            "john.doe@example.com",
		"password":         "password123",
		"confirm_password": "password123",
		"city":             "北京",
		"address": map[string]any{
			"city":   "Beijing",
			"street": "Long St",
		},
		"ptr_address": map[string]any{
			"city":   "上海",
			"street": "花山路",
		},
	}

	testCases := []struct {
		name     string
		input    any
		expected map[string]any
		wantErr  bool
	}{
		{
			name:     "Basic Struct",
			input:    u,
			expected: expectedMap,
			wantErr:  false,
		},
		{
			name:     "Pointer to Struct",
			input:    &u,
			expected: expectedMap,
			wantErr:  false,
		},
		{
			name:     "Nil Pointer",
			input:    (*User)(nil),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Non-Struct Input",
			input:    "not a struct",
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Struct with Nil Pointer Field",
			input: User{
				Name:            "Jane",
				Age:             nil, //  空指针字段
				Email:           "jane@example.com",
				Password:        "password456",
				ConfirmPassword: "password456",
				City:            "Shenzhen",
				Address:         Address{City: "Shenzhen", Street: "Short St"},
			},
			expected: map[string]any{
				"name":             "Jane",
				"age":              nil, // 空指针字段的值为 nil
				"email":            "jane@example.com",
				"password":         "password456",
				"confirm_password": "password456",
				"city":             "Shenzhen",
				"address":          map[string]any{"city": "Shenzhen", "street": "Short St"},
				"ptr_address":      nil,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualMap, err := structToMap(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("structToMap() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(actualMap, tc.expected) {
				t.Errorf("structToMap() = %v, want %v", actualMap, tc.expected)
			}
		})
	}
}
