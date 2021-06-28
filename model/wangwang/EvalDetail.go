package wangwang

// EvalDetail 
/* model for simplify = false
type EvalDetail struct {

    // 评分：0-非常满意；1-满意；2-一般；3-不满意；4-非常不满意
    
    EvalCode   int64 `json:"eval_code,omitempty"`
    

    // 接收评价的消费者昵称
    
    EvalRecer   string `json:"eval_recer,omitempty"`
    

    // 发送评价邀请的商家客服昵称
    
    EvalSender   string `json:"eval_sender,omitempty"`
    

    // 最后一次评价的时间
    
    EvalTime   string `json:"eval_time,omitempty"`
    

    // 评价的发送时间
    
    SendTime   string `json:"send_time,omitempty"`
    

    // 评价来源：0-客服邀评；1-消费者自主评价；2-系统邀评
    
    Source   int64 `json:"source,omitempty"`
    

}
*/

// EvalDetail 
type EvalDetail struct {

    // 评分：0-非常满意；1-满意；2-一般；3-不满意；4-非常不满意
    EvalCode   int64 `json:"eval_code,omitempty"`

    // 接收评价的消费者昵称
    EvalRecer   string `json:"eval_recer,omitempty"`

    // 发送评价邀请的商家客服昵称
    EvalSender   string `json:"eval_sender,omitempty"`

    // 最后一次评价的时间
    EvalTime   string `json:"eval_time,omitempty"`

    // 评价的发送时间
    SendTime   string `json:"send_time,omitempty"`

    // 评价来源：0-客服邀评；1-消费者自主评价；2-系统邀评
    Source   int64 `json:"source,omitempty"`

}
