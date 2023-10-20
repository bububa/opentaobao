package alihouse

import (
	"sync"
)

// IMReceiveModelReqDto 结构体
type IMReceiveModelReqDto struct {
	// 外部id列表，根据type区分
	TagIds []string `json:"tag_ids,omitempty" xml:"tag_ids>string,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 0-楼盘 1-交易商品
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 1-天猫好房来客通 2-千牛
	ImModel int64 `json:"im_model,omitempty" xml:"im_model,omitempty"`
	// 0-非测试 1-测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolIMReceiveModelReqDto = sync.Pool{
	New: func() any {
		return new(IMReceiveModelReqDto)
	},
}

// GetIMReceiveModelReqDto() 从对象池中获取IMReceiveModelReqDto
func GetIMReceiveModelReqDto() *IMReceiveModelReqDto {
	return poolIMReceiveModelReqDto.Get().(*IMReceiveModelReqDto)
}

// ReleaseIMReceiveModelReqDto 释放IMReceiveModelReqDto
func ReleaseIMReceiveModelReqDto(v *IMReceiveModelReqDto) {
	v.TagIds = v.TagIds[:0]
	v.OuterStoreId = ""
	v.Type = 0
	v.ImModel = 0
	v.IsTest = 0
	poolIMReceiveModelReqDto.Put(v)
}
