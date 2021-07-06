package fenxiao

// TopMemoDto 结构体
type TopMemoDto struct {
	// attachments
	Attachments []TopMemoAttachment `json:"attachments,omitempty" xml:"attachments>top_memo_attachment,omitempty"`
	// remark
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// operateUserNick
	OperateUserNick string `json:"operate_user_nick,omitempty" xml:"operate_user_nick,omitempty"`
}
