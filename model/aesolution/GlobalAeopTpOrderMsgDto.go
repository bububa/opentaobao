package aesolution

// GlobalAeopTpOrderMsgDto 结构体
type GlobalAeopTpOrderMsgDto struct {
	// order id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// order creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// order modification time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// business order id
	BusinessOrderId int64 `json:"business_order_id,omitempty" xml:"business_order_id,omitempty"`
	// message status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// meesage content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// message sender ( buyer; seller)
	Poster string `json:"poster,omitempty" xml:"poster,omitempty"`
	// sender login ID
	SenderLoginId string `json:"sender_login_id,omitempty" xml:"sender_login_id,omitempty"`
	// sender account seq
	SenderSeq int64 `json:"sender_seq,omitempty" xml:"sender_seq,omitempty"`
	// sender admin account seq
	SenderAdminSeq int64 `json:"sender_admin_seq,omitempty" xml:"sender_admin_seq,omitempty"`
	// sender first name
	SenderFirstName string `json:"sender_first_name,omitempty" xml:"sender_first_name,omitempty"`
	// send last name
	SenderLastName string `json:"sender_last_name,omitempty" xml:"sender_last_name,omitempty"`
	// receiver ID
	ReceiverLoginId string `json:"receiver_login_id,omitempty" xml:"receiver_login_id,omitempty"`
	// receiver account seq
	ReceiverSeq int64 `json:"receiver_seq,omitempty" xml:"receiver_seq,omitempty"`
	// receiver admin account seq
	ReceiverAdminSeq int64 `json:"receiver_admin_seq,omitempty" xml:"receiver_admin_seq,omitempty"`
	// receiver first name
	ReceiverFirstName string `json:"receiver_first_name,omitempty" xml:"receiver_first_name,omitempty"`
	// receiver last name
	ReceiverLastName string `json:"receiver_last_name,omitempty" xml:"receiver_last_name,omitempty"`
}
