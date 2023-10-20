package idle

import (
	"sync"
)

// IdleAppraiseSpuRegister4TopDto 结构体
type IdleAppraiseSpuRegister4TopDto struct {
	// 类目Id
	CateId string `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
	// 操作类型，0新增，-1删除。当spu第一次挂载时，会进入1测试中状态。服务商联调通过后，需要再次挂载，actionType还传0，挂载信息状态会变成0已上线。
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 验货类型，1新品，2二手
	AppraiseScene int64 `json:"appraise_scene,omitempty" xml:"appraise_scene,omitempty"`
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// spuId
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolIdleAppraiseSpuRegister4TopDto = sync.Pool{
	New: func() any {
		return new(IdleAppraiseSpuRegister4TopDto)
	},
}

// GetIdleAppraiseSpuRegister4TopDto() 从对象池中获取IdleAppraiseSpuRegister4TopDto
func GetIdleAppraiseSpuRegister4TopDto() *IdleAppraiseSpuRegister4TopDto {
	return poolIdleAppraiseSpuRegister4TopDto.Get().(*IdleAppraiseSpuRegister4TopDto)
}

// ReleaseIdleAppraiseSpuRegister4TopDto 释放IdleAppraiseSpuRegister4TopDto
func ReleaseIdleAppraiseSpuRegister4TopDto(v *IdleAppraiseSpuRegister4TopDto) {
	v.CateId = ""
	v.ActionType = 0
	v.AppraiseScene = 0
	v.BrandId = 0
	v.SpuId = 0
	poolIdleAppraiseSpuRegister4TopDto.Put(v)
}
