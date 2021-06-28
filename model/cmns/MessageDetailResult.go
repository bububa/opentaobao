package cmns

// MessageDetailResult 
type MessageDetailResult struct {

    // 消息到达数
    
    AckCnt   int64 `json:"ack_cnt,omitempty" xml:"ack_cnt,omitempty"`
    

    // 消息侦听通道
    
    Action   string `json:"action,omitempty" xml:"action,omitempty"`
    

    // appKey
    
    AppKey   string `json:"app_key,omitempty" xml:"app_key,omitempty"`
    

    // 应用名称
    
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    

    // 消息审核状态
    
    Audit   int64 `json:"audit,omitempty" xml:"audit,omitempty"`
    

    // 通知消息内容
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 消息失效时间
    
    ExipiredTime   string `json:"exipired_time,omitempty" xml:"exipired_time,omitempty"`
    

    // 消息创建时间
    
    GmtCreateTime   string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
    

    // 消息id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 消息参数
    
    Parameter   string `json:"parameter,omitempty" xml:"parameter,omitempty"`
    

    // 预计发送数
    
    PredictSendCnt   int64 `json:"predict_send_cnt,omitempty" xml:"predict_send_cnt,omitempty"`
    

    // 消息优先级
    
    Priority   int64 `json:"priority,omitempty" xml:"priority,omitempty"`
    

    // 消息接收者json string
    
    Receiver   string `json:"receiver,omitempty" xml:"receiver,omitempty"`
    

    // 消息发送状态 0:发送中，1:发送完成，2:发送过期
    
    SendStatus   int64 `json:"send_status,omitempty" xml:"send_status,omitempty"`
    

    // 通知消息标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 消息类型，1为透传，2为通知
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 消息uri,yunos4.x系统专用
    
    Uri   string `json:"uri,omitempty" xml:"uri,omitempty"`
    

}
