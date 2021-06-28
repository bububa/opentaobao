package legalcase

// NoticeModel 
/* model for simplify = false
type NoticeModel struct {

    // 创建人
    
    Creator   string `json:"creator,omitempty"`
    

    // 附件
    
    Attachment  struct {
        FileModel  []FileModel `json:"file_model,omitempty"`
    } `json:"attachment,omitempty"`
    

    // 消息唯一id
    
    NoticeId   string `json:"notice_id,omitempty"`
    

    // 内容
    
    Content   string `json:"content,omitempty"`
    

}
*/

// NoticeModel 
type NoticeModel struct {

    // 创建人
    Creator   string `json:"creator,omitempty"`

    // 附件
    Attachment   []FileModel `json:"attachment,omitempty"`

    // 消息唯一id
    NoticeId   string `json:"notice_id,omitempty"`

    // 内容
    Content   string `json:"content,omitempty"`

}
