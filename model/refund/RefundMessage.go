package refund

// RefundMessage 
type RefundMessage struct {
    // 留言编号
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 留言创建时间。格式:yyyy-MM-dd HH:mm:ss
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 留言者编号
    OwnerId   int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
    // 留言者昵称
    OwnerNick   string `json:"owner_nick,omitempty" xml:"owner_nick,omitempty"`
    // 留言内容。最大长度: 400个字节
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 凭证附件地址（图片）
    PicUrls   []PicUrl `json:"pic_urls,omitempty" xml:"pic_urls>pic_url,omitempty"`
    // 退款类型：NORMAL（普通留言），RETURN_GOODS_APPROVED（卖家留退货地址时留言）；如果为RETURN_GOODS_APPROVED，则退款留言中有卖家收货地址
    MessageType   string `json:"message_type,omitempty" xml:"message_type,omitempty"`
    // 退款阶段，可选值：onsale（售中）, aftersale（售后）
    RefundPhase   string `json:"refund_phase,omitempty" xml:"refund_phase,omitempty"`
    // 退款编号。
    RefundId   string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    // 留言者身份1代表买家，2代表卖家，3代表小二
    OwnerRole   string `json:"owner_role,omitempty" xml:"owner_role,omitempty"`
}
