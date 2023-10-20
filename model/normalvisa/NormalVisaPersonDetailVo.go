package normalvisa

import (
	"sync"
)

// NormalVisaPersonDetailVo 结构体
type NormalVisaPersonDetailVo struct {
	// 证件号
	CredentialCardNo string `json:"credential_card_no,omitempty" xml:"credential_card_no,omitempty"`
	// 当前状态：1001,1002,1003,1004,1005,1006,1007,1008,1010
	CurrentStatusDesc string `json:"current_status_desc,omitempty" xml:"current_status_desc,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 中止状态
	FinishStatusDesc string `json:"finish_status_desc,omitempty" xml:"finish_status_desc,omitempty"`
	// 下一个状态描述：无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
	NextCurrentStatusDesc string `json:"next_current_status_desc,omitempty" xml:"next_current_status_desc,omitempty"`
	// 证件信息
	CredentialCardInfor string `json:"credential_card_infor,omitempty" xml:"credential_card_infor,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	PersonVisaId int64 `json:"person_visa_id,omitempty" xml:"person_visa_id,omitempty"`
	// 当前状态：无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
	CurrentStatus int64 `json:"current_status,omitempty" xml:"current_status,omitempty"`
	// 中止状态
	FinishStatus int64 `json:"finish_status,omitempty" xml:"finish_status,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 下一个状态：无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
	NextCurrentStatus int64 `json:"next_current_status,omitempty" xml:"next_current_status,omitempty"`
	// 是否出签，为空则买家没有反馈，1 表示出签，0 表示拒签，-1 表示未办理
	Pass int64 `json:"pass,omitempty" xml:"pass,omitempty"`
	// 是否有下个状态
	HasNextStatus bool `json:"has_next_status,omitempty" xml:"has_next_status,omitempty"`
	// 是否包含结束状态
	HasFinishStatus bool `json:"has_finish_status,omitempty" xml:"has_finish_status,omitempty"`
	// 是否达到最终状态
	Disabled bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
}

var poolNormalVisaPersonDetailVo = sync.Pool{
	New: func() any {
		return new(NormalVisaPersonDetailVo)
	},
}

// GetNormalVisaPersonDetailVo() 从对象池中获取NormalVisaPersonDetailVo
func GetNormalVisaPersonDetailVo() *NormalVisaPersonDetailVo {
	return poolNormalVisaPersonDetailVo.Get().(*NormalVisaPersonDetailVo)
}

// ReleaseNormalVisaPersonDetailVo 释放NormalVisaPersonDetailVo
func ReleaseNormalVisaPersonDetailVo(v *NormalVisaPersonDetailVo) {
	v.CredentialCardNo = ""
	v.CurrentStatusDesc = ""
	v.Remark = ""
	v.FinishStatusDesc = ""
	v.NextCurrentStatusDesc = ""
	v.CredentialCardInfor = ""
	v.Name = ""
	v.PersonVisaId = 0
	v.CurrentStatus = 0
	v.FinishStatus = 0
	v.BizOrderId = 0
	v.NextCurrentStatus = 0
	v.Pass = 0
	v.HasNextStatus = false
	v.HasFinishStatus = false
	v.Disabled = false
	poolNormalVisaPersonDetailVo.Put(v)
}
