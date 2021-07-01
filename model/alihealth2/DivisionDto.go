package alihealth2

// DivisionDto 结构体
type DivisionDto struct {
	// divisionName
	DivisionName string `json:"division_name,omitempty" xml:"division_name,omitempty"`
	// divisionId
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}
