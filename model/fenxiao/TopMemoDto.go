package fenxiao

// TopMemoDto 
/* model for simplify = false
type TopMemoDto struct {

    // remark
    
    Remark   string `json:"remark,omitempty"`
    

    // operateUserNick
    
    OperateUserNick   string `json:"operate_user_nick,omitempty"`
    

    // attachments
    
    Attachments  struct {
        TopMemoAttachment  []TopMemoAttachment `json:"top_memo_attachment,omitempty"`
    } `json:"attachments,omitempty"`
    

}
*/

// TopMemoDto 
type TopMemoDto struct {

    // remark
    Remark   string `json:"remark,omitempty"`

    // operateUserNick
    OperateUserNick   string `json:"operate_user_nick,omitempty"`

    // attachments
    Attachments   []TopMemoAttachment `json:"attachments,omitempty"`

}
