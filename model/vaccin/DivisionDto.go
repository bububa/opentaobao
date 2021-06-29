package vaccin

// DivisionDTO 
type DivisionDTO struct {
    // childDivision
    ChildDivisionList   []DivisionDTO `json:"child_division_list,omitempty" xml:"child_division_list>division_dto,omitempty"`
    // divisionId
    DivisionId   int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
    // divisionName
    DivisionName   string `json:"division_name,omitempty" xml:"division_name,omitempty"`
}
