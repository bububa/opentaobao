package fenxiao

// OrderMessages 
type OrderMessages struct {

    // 留言时间
    MessageTime   string `json:"message_time,omitempty"`

    // 留言标题，例如：分销商留言，供应商留言，买家留言
    MessageTitle   string `json:"message_title,omitempty"`

    // 留言内容
    MessageContent   string `json:"message_content,omitempty"`

    // 留言时的图片地址
    PicUrl   string `json:"pic_url,omitempty"`

}
