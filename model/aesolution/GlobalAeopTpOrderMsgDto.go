package aesolution

import (
	"sync"
)

// GlobalAeopTpOrderMsgDto 结构体
type GlobalAeopTpOrderMsgDto struct {
	// order creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// order modification time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// message status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// meesage content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// message sender ( buyer; seller)
	Poster string `json:"poster,omitempty" xml:"poster,omitempty"`
	// sender login ID
	SenderLoginId string `json:"sender_login_id,omitempty" xml:"sender_login_id,omitempty"`
	// sender first name
	SenderFirstName string `json:"sender_first_name,omitempty" xml:"sender_first_name,omitempty"`
	// send last name
	SenderLastName string `json:"sender_last_name,omitempty" xml:"sender_last_name,omitempty"`
	// receiver ID
	ReceiverLoginId string `json:"receiver_login_id,omitempty" xml:"receiver_login_id,omitempty"`
	// receiver first name
	ReceiverFirstName string `json:"receiver_first_name,omitempty" xml:"receiver_first_name,omitempty"`
	// receiver last name
	ReceiverLastName string `json:"receiver_last_name,omitempty" xml:"receiver_last_name,omitempty"`
	// order id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// business order id
	BusinessOrderId int64 `json:"business_order_id,omitempty" xml:"business_order_id,omitempty"`
	// sender account seq
	SenderSeq int64 `json:"sender_seq,omitempty" xml:"sender_seq,omitempty"`
	// sender admin account seq
	SenderAdminSeq int64 `json:"sender_admin_seq,omitempty" xml:"sender_admin_seq,omitempty"`
	// receiver account seq
	ReceiverSeq int64 `json:"receiver_seq,omitempty" xml:"receiver_seq,omitempty"`
	// receiver admin account seq
	ReceiverAdminSeq int64 `json:"receiver_admin_seq,omitempty" xml:"receiver_admin_seq,omitempty"`
}

var poolGlobalAeopTpOrderMsgDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpOrderMsgDto)
	},
}

// GetGlobalAeopTpOrderMsgDto() 从对象池中获取GlobalAeopTpOrderMsgDto
func GetGlobalAeopTpOrderMsgDto() *GlobalAeopTpOrderMsgDto {
	return poolGlobalAeopTpOrderMsgDto.Get().(*GlobalAeopTpOrderMsgDto)
}

// ReleaseGlobalAeopTpOrderMsgDto 释放GlobalAeopTpOrderMsgDto
func ReleaseGlobalAeopTpOrderMsgDto(v *GlobalAeopTpOrderMsgDto) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Status = ""
	v.Content = ""
	v.Poster = ""
	v.SenderLoginId = ""
	v.SenderFirstName = ""
	v.SenderLastName = ""
	v.ReceiverLoginId = ""
	v.ReceiverFirstName = ""
	v.ReceiverLastName = ""
	v.Id = 0
	v.BusinessOrderId = 0
	v.SenderSeq = 0
	v.SenderAdminSeq = 0
	v.ReceiverSeq = 0
	v.ReceiverAdminSeq = 0
	poolGlobalAeopTpOrderMsgDto.Put(v)
}
