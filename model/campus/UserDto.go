package campus

import (
	"sync"
)

// UserDto 结构体
type UserDto struct {
	// 用户自定义ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}

var poolUserDto = sync.Pool{
	New: func() any {
		return new(UserDto)
	},
}

// GetUserDto() 从对象池中获取UserDto
func GetUserDto() *UserDto {
	return poolUserDto.Get().(*UserDto)
}

// ReleaseUserDto 释放UserDto
func ReleaseUserDto(v *UserDto) {
	v.UserId = ""
	v.Name = ""
	v.CompanyId = 0
	poolUserDto.Put(v)
}
