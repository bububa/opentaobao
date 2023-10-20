package alihealthcrm

import (
	"sync"
)

// BatchRemindRequestDto 结构体
type BatchRemindRequestDto struct {
	// 模板参数，可选,|分割
	Params string `json:"params,omitempty" xml:"params,omitempty"`
	// 访问跳转链接
	VisitUrl string `json:"visit_url,omitempty" xml:"visit_url,omitempty"`
	// 提醒宝宝ID
	BabyId int64 `json:"baby_id,omitempty" xml:"baby_id,omitempty"`
	// 用户健康ID
	TpUserId int64 `json:"tp_user_id,omitempty" xml:"tp_user_id,omitempty"`
}

var poolBatchRemindRequestDto = sync.Pool{
	New: func() any {
		return new(BatchRemindRequestDto)
	},
}

// GetBatchRemindRequestDto() 从对象池中获取BatchRemindRequestDto
func GetBatchRemindRequestDto() *BatchRemindRequestDto {
	return poolBatchRemindRequestDto.Get().(*BatchRemindRequestDto)
}

// ReleaseBatchRemindRequestDto 释放BatchRemindRequestDto
func ReleaseBatchRemindRequestDto(v *BatchRemindRequestDto) {
	v.Params = ""
	v.VisitUrl = ""
	v.BabyId = 0
	v.TpUserId = 0
	poolBatchRemindRequestDto.Put(v)
}
