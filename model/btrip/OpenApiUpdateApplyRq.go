package btrip

import (
	"sync"
)

// OpenApiUpdateApplyRq 结构体
type OpenApiUpdateApplyRq struct {
	// 外部申请单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 备注
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 审批人id（第三方用户Id）
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审批人名字
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 1已同意 2已拒绝 3已转交 4已取消
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 1、老版本2、isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenApiUpdateApplyRq = sync.Pool{
	New: func() any {
		return new(OpenApiUpdateApplyRq)
	},
}

// GetOpenApiUpdateApplyRq() 从对象池中获取OpenApiUpdateApplyRq
func GetOpenApiUpdateApplyRq() *OpenApiUpdateApplyRq {
	return poolOpenApiUpdateApplyRq.Get().(*OpenApiUpdateApplyRq)
}

// ReleaseOpenApiUpdateApplyRq 释放OpenApiUpdateApplyRq
func ReleaseOpenApiUpdateApplyRq(v *OpenApiUpdateApplyRq) {
	v.ThirdpartApplyId = ""
	v.CorpId = ""
	v.Note = ""
	v.OperateTime = ""
	v.UserId = ""
	v.UserName = ""
	v.Status = 0
	v.Version = 0
	poolOpenApiUpdateApplyRq.Put(v)
}
