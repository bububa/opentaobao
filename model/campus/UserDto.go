package campus

// UserDto 结构体
type UserDto struct {
	// 用户自定义ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}
