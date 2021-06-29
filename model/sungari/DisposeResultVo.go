package sungari

// DisposeResultVo 
type DisposeResultVo struct {
    // 完成时间
    FinishTime   string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
    // 结果描述
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 状态描述
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 回复函保存的ossKey，需要自行下载
    ReplyKey   string `json:"reply_key,omitempty" xml:"reply_key,omitempty"`
    // 创建人
    CreatePerson   string `json:"create_person,omitempty" xml:"create_person,omitempty"`
    // id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 回复函链接
    ReplyUrl   string `json:"reply_url,omitempty" xml:"reply_url,omitempty"`
}
