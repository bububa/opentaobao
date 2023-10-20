package wangwang

import (
	"sync"
)

// EvalDetail 结构体
type EvalDetail struct {
	// 发送评价邀请的商家客服昵称
	EvalSender string `json:"eval_sender,omitempty" xml:"eval_sender,omitempty"`
	// 最后一次评价的时间
	EvalTime string `json:"eval_time,omitempty" xml:"eval_time,omitempty"`
	// 接收评价的消费者用户ID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 评价的发送时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 评价标签，可空
	LabelName string `json:"label_name,omitempty" xml:"label_name,omitempty"`
	// 脱敏后的买家nick，可空
	EvalRecer string `json:"eval_recer,omitempty" xml:"eval_recer,omitempty"`
	// 评分：0-非常满意；1-满意；2-一般；3-不满意；4-非常不满意
	EvalCode int64 `json:"eval_code,omitempty" xml:"eval_code,omitempty"`
	// 评价来源：0-客服邀评；1-消费者自主评价；2-系统邀评
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
}

var poolEvalDetail = sync.Pool{
	New: func() any {
		return new(EvalDetail)
	},
}

// GetEvalDetail() 从对象池中获取EvalDetail
func GetEvalDetail() *EvalDetail {
	return poolEvalDetail.Get().(*EvalDetail)
}

// ReleaseEvalDetail 释放EvalDetail
func ReleaseEvalDetail(v *EvalDetail) {
	v.EvalSender = ""
	v.EvalTime = ""
	v.OpenUid = ""
	v.SendTime = ""
	v.LabelName = ""
	v.EvalRecer = ""
	v.EvalCode = 0
	v.Source = 0
	poolEvalDetail.Put(v)
}
