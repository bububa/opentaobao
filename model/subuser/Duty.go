package subuser

// Duty 
type Duty struct {
    // 职务ID
    DutyId   int64 `json:"duty_id,omitempty" xml:"duty_id,omitempty"`
    // 职务名称
    DutyName   string `json:"duty_name,omitempty" xml:"duty_name,omitempty"`
    // 职务级别
    DutyLevel   int64 `json:"duty_level,omitempty" xml:"duty_level,omitempty"`
}
