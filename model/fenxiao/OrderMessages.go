package fenxiao

import (
	"sync"
)

// OrderMessages 结构体
type OrderMessages struct {
	// 留言时间
	MessageTime string `json:"message_time,omitempty" xml:"message_time,omitempty"`
	// 留言标题，例如：分销商留言，供应商留言，买家留言
	MessageTitle string `json:"message_title,omitempty" xml:"message_title,omitempty"`
	// 留言内容
	MessageContent string `json:"message_content,omitempty" xml:"message_content,omitempty"`
	// 留言时的图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
}

var poolOrderMessages = sync.Pool{
	New: func() any {
		return new(OrderMessages)
	},
}

// GetOrderMessages() 从对象池中获取OrderMessages
func GetOrderMessages() *OrderMessages {
	return poolOrderMessages.Get().(*OrderMessages)
}

// ReleaseOrderMessages 释放OrderMessages
func ReleaseOrderMessages(v *OrderMessages) {
	v.MessageTime = ""
	v.MessageTitle = ""
	v.MessageContent = ""
	v.PicUrl = ""
	poolOrderMessages.Put(v)
}
