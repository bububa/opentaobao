package openmall

import (
	"sync"
)

// RefundMessage 结构体
type RefundMessage struct {
	// 退款单结构
	RefundMessagePics []RefundMessagePic `json:"refund_message_pics,omitempty" xml:"refund_message_pics>refund_message_pic,omitempty"`
	// 退款留言
	RefundMessage string `json:"refund_message,omitempty" xml:"refund_message,omitempty"`
	// 留言时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 操作角色ID，1（卖家主账户），2（卖家子账户），3（小二）、4（买家）、5（系统）、6（系统超时）、7（服务商）；openmall中提交的留言角色均为4买家身份
	Role int64 `json:"role,omitempty" xml:"role,omitempty"`
}

var poolRefundMessage = sync.Pool{
	New: func() any {
		return new(RefundMessage)
	},
}

// GetRefundMessage() 从对象池中获取RefundMessage
func GetRefundMessage() *RefundMessage {
	return poolRefundMessage.Get().(*RefundMessage)
}

// ReleaseRefundMessage 释放RefundMessage
func ReleaseRefundMessage(v *RefundMessage) {
	v.RefundMessagePics = v.RefundMessagePics[:0]
	v.RefundMessage = ""
	v.Time = ""
	v.Role = 0
	poolRefundMessage.Put(v)
}
