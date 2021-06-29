package exchange

// RefundMessage 
type RefundMessage struct {
    // 换货单号ID
    RefundId   int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    // 留言创建时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 留言ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 留言者ID
    OwnerId   int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
    // 留言者昵称
    OwnerNick   string `json:"owner_nick,omitempty" xml:"owner_nick,omitempty"`
    // 留言内容
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 凭证信息
    PicUrls   []PicUrl `json:"pic_urls,omitempty" xml:"pic_urls>pic_url,omitempty"`
    // 留言类型：系统或是人工
    MessageType   string `json:"message_type,omitempty" xml:"message_type,omitempty"`
    // 留言者橘色
    OwnerRole   string `json:"owner_role,omitempty" xml:"owner_role,omitempty"`
}
