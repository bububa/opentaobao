package jipiao

import (
	"sync"
)

// RequestAxbDo 结构体
type RequestAxbDo struct {
	// 商家联系号码,多个号码以“,”分割；传证件号时为乘机人姓名
	ContactNo string `json:"contact_no,omitempty" xml:"contact_no,omitempty"`
	// 用途
	Purpose string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// 请求内容：0手机号
	ReqContent int64 `json:"req_content,omitempty" xml:"req_content,omitempty"`
	// 业务类型：0国内机票,1国际机票
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolRequestAxbDo = sync.Pool{
	New: func() any {
		return new(RequestAxbDo)
	},
}

// GetRequestAxbDo() 从对象池中获取RequestAxbDo
func GetRequestAxbDo() *RequestAxbDo {
	return poolRequestAxbDo.Get().(*RequestAxbDo)
}

// ReleaseRequestAxbDo 释放RequestAxbDo
func ReleaseRequestAxbDo(v *RequestAxbDo) {
	v.ContactNo = ""
	v.Purpose = ""
	v.ReqContent = 0
	v.BizType = 0
	v.OrderId = 0
	poolRequestAxbDo.Put(v)
}
