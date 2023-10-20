package cainiaohandover

import (
	"sync"
)

// HandoverContentUpdateParcelDto 结构体
type HandoverContentUpdateParcelDto struct {
	// 小包对应的店铺账号;比如cnxxxx;填入补充相关信息性能更好
	LoginId string `json:"login_id,omitempty" xml:"login_id,omitempty"`
	// 小包的LP号,必填;
	LpCode string `json:"lp_code,omitempty" xml:"lp_code,omitempty"`
	// 小包对应的店铺id;填入相关信息性能更好
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolHandoverContentUpdateParcelDto = sync.Pool{
	New: func() any {
		return new(HandoverContentUpdateParcelDto)
	},
}

// GetHandoverContentUpdateParcelDto() 从对象池中获取HandoverContentUpdateParcelDto
func GetHandoverContentUpdateParcelDto() *HandoverContentUpdateParcelDto {
	return poolHandoverContentUpdateParcelDto.Get().(*HandoverContentUpdateParcelDto)
}

// ReleaseHandoverContentUpdateParcelDto 释放HandoverContentUpdateParcelDto
func ReleaseHandoverContentUpdateParcelDto(v *HandoverContentUpdateParcelDto) {
	v.LoginId = ""
	v.LpCode = ""
	v.SellerId = 0
	poolHandoverContentUpdateParcelDto.Put(v)
}
