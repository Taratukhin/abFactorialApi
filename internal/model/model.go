package model
package model

type CheckDataType struct { // this structure is needed to check the presence of these fields in the json
	A *int `json:"a"`
	B *int `json:"b"`
}
type DataType struct { // this structure is needed for the result
	A uint64 `json:"a"`
	B uint64 `json:"b"`
}

