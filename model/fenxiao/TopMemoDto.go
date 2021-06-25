package fenxiao

// TopMemoDto 
type TopMemoDto struct {

    // remark
    Remark   string `json:"remark,omitempty"`

    // operateUserNick
    OperateUserNick   string `json:"operate_user_nick,omitempty"`

    // attachments
    Attachments   []TopMemoAttachment `json:"attachments,omitempty"`

}
