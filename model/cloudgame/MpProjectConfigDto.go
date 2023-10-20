package cloudgame

import (
	"sync"
)

// MpProjectConfigDto 结构体
type MpProjectConfigDto struct {
	// user在mp上的accountid
	ConfigKey string `json:"config_key,omitempty" xml:"config_key,omitempty"`
	// user的acesstoken
	CustomerUniqueId string `json:"customer_unique_id,omitempty" xml:"customer_unique_id,omitempty"`
	// customer platform env
	CustomeEnv string `json:"custome_env,omitempty" xml:"custome_env,omitempty"`
	// customer project id
	CustomerProjectId string `json:"customer_project_id,omitempty" xml:"customer_project_id,omitempty"`
	// customer id
	CustomeId int64 `json:"custome_id,omitempty" xml:"custome_id,omitempty"`
	// 是否检查userToken，目前都是0
	CheckUserToken int64 `json:"check_user_token,omitempty" xml:"check_user_token,omitempty"`
	// 是否只有一个block，目前都是1
	OnlyOneBlock int64 `json:"only_one_block,omitempty" xml:"only_one_block,omitempty"`
	// 缺省的blockid
	DefaultMpBlockId int64 `json:"default_mp_block_id,omitempty" xml:"default_mp_block_id,omitempty"`
	// mpprojectid
	MpProjectId int64 `json:"mp_project_id,omitempty" xml:"mp_project_id,omitempty"`
}

var poolMpProjectConfigDto = sync.Pool{
	New: func() any {
		return new(MpProjectConfigDto)
	},
}

// GetMpProjectConfigDto() 从对象池中获取MpProjectConfigDto
func GetMpProjectConfigDto() *MpProjectConfigDto {
	return poolMpProjectConfigDto.Get().(*MpProjectConfigDto)
}

// ReleaseMpProjectConfigDto 释放MpProjectConfigDto
func ReleaseMpProjectConfigDto(v *MpProjectConfigDto) {
	v.ConfigKey = ""
	v.CustomerUniqueId = ""
	v.CustomeEnv = ""
	v.CustomerProjectId = ""
	v.CustomeId = 0
	v.CheckUserToken = 0
	v.OnlyOneBlock = 0
	v.DefaultMpBlockId = 0
	v.MpProjectId = 0
	poolMpProjectConfigDto.Put(v)
}
