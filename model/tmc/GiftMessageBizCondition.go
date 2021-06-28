package tmc

// GiftMessageBizCondition 
type GiftMessageBizCondition struct {

    // 分页参数
    
    Paginator   *Paginator `json:"paginator,omitempty" xml:"paginator,omitempty"`
    

    // 有效时间
    
    ValidateDate   string `json:"validate_date,omitempty" xml:"validate_date,omitempty"`
    

    // 系统自动生成
    
    ReceiverId   string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
    

    // 消息状态
    
    MessageStatus   int64 `json:"message_status,omitempty" xml:"message_status,omitempty"`
    

    // 消息id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 按时间排序1:asc 2:desc
    
    Sort   int64 `json:"sort,omitempty" xml:"sort,omitempty"`
    

}
