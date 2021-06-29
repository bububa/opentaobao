package drugtrace

// AlibabaAlihealthDrugCodeKytYyApplycodeModel 
type AlibabaAlihealthDrugCodeKytYyApplycodeModel struct {
    // 子码
    ChildrenCodeList   []string `json:"children_code_list,omitempty" xml:"children_code_list>string,omitempty"`
    // 父码
    ParentCode   string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}
