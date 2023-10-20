package flight

import (
	"sync"
)

// CaseResultDetailDto 结构体
type CaseResultDetailDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 协同单标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 协同单截止时间
	ServeDeadline string `json:"serve_deadline,omitempty" xml:"serve_deadline,omitempty"`
	// 飞猪订单号
	CorrelationOutOrderId string `json:"correlation_out_order_id,omitempty" xml:"correlation_out_order_id,omitempty"`
	// 创建人
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 跟进人
	FollowerName string `json:"follower_name,omitempty" xml:"follower_name,omitempty"`
	// 协同请求内容
	Request string `json:"request,omitempty" xml:"request,omitempty"`
	// 协同回复内容
	Reply string `json:"reply,omitempty" xml:"reply,omitempty"`
	// 表单内容
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 完结时间
	OperateFinishTime string `json:"operate_finish_time,omitempty" xml:"operate_finish_time,omitempty"`
	// 催促人姓名
	UrgerName string `json:"urger_name,omitempty" xml:"urger_name,omitempty"`
	// 协同单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 问题类型
	CaseType int64 `json:"case_type,omitempty" xml:"case_type,omitempty"`
	// 协同单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否催促标识
	Urge int64 `json:"urge,omitempty" xml:"urge,omitempty"`
	// 关联业务类型
	CorrelationBizType int64 `json:"correlation_biz_type,omitempty" xml:"correlation_biz_type,omitempty"`
	// 跟进人时长(时间戳)
	FollowTime int64 `json:"follow_time,omitempty" xml:"follow_time,omitempty"`
	// 支付状态(0:未支付,1:已支付,2:支付取消,3:无需支付)
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 是否需要支付
	NeedPayFlag bool `json:"need_pay_flag,omitempty" xml:"need_pay_flag,omitempty"`
}

var poolCaseResultDetailDto = sync.Pool{
	New: func() any {
		return new(CaseResultDetailDto)
	},
}

// GetCaseResultDetailDto() 从对象池中获取CaseResultDetailDto
func GetCaseResultDetailDto() *CaseResultDetailDto {
	return poolCaseResultDetailDto.Get().(*CaseResultDetailDto)
}

// ReleaseCaseResultDetailDto 释放CaseResultDetailDto
func ReleaseCaseResultDetailDto(v *CaseResultDetailDto) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Title = ""
	v.ServeDeadline = ""
	v.CorrelationOutOrderId = ""
	v.CreatorName = ""
	v.FollowerName = ""
	v.Request = ""
	v.Reply = ""
	v.ExtraInfo = ""
	v.EndTime = ""
	v.OperateFinishTime = ""
	v.UrgerName = ""
	v.Id = 0
	v.CaseType = 0
	v.Status = 0
	v.Urge = 0
	v.CorrelationBizType = 0
	v.FollowTime = 0
	v.PayStatus = 0
	v.NeedPayFlag = false
	poolCaseResultDetailDto.Put(v)
}
