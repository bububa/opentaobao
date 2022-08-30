package fenxiao

// TopMemoDto 结构体
type TopMemoDto struct {
	// 附件
	Attachments []TopMemoAttachment `json:"attachments,omitempty" xml:"attachments>top_memo_attachment,omitempty"`
	// 子订单备注内容
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作者昵称
	OperateUserNick string `json:"operate_user_nick,omitempty" xml:"operate_user_nick,omitempty"`
}
