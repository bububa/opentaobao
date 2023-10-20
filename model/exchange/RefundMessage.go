package exchange

import (
	"sync"
)

// RefundMessage 结构体
type RefundMessage struct {
	// 凭证信息
	PicUrls []PicUrl `json:"pic_urls,omitempty" xml:"pic_urls>pic_url,omitempty"`
	// 留言者昵称
	OwnerNick string `json:"owner_nick,omitempty" xml:"owner_nick,omitempty"`
	// 留言内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 留言创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 留言类型：系统或是人工
	MessageType string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	// 留言者橘色
	OwnerRole string `json:"owner_role,omitempty" xml:"owner_role,omitempty"`
	// 留言者openId
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 留言ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 换货单号ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 留言者ID
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
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
	v.PicUrls = v.PicUrls[:0]
	v.OwnerNick = ""
	v.Content = ""
	v.Created = ""
	v.MessageType = ""
	v.OwnerRole = ""
	v.OpenUid = ""
	v.Id = 0
	v.RefundId = 0
	v.OwnerId = 0
	poolRefundMessage.Put(v)
}
