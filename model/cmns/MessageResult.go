package cmns

// MessageResult 
/* model for simplify = false
type MessageResult struct {

    // 消息过期时间
    
    ExpireTime   int64 `json:"expire_time,omitempty"`
    

    // 消息预计发送数
    
    Total2send   int64 `json:"total2send,omitempty"`
    

    // 消息审核状态1,3:待审核 2:审核通过，处于发送中 4:审核拒绝 5 已撤回 6 已去重
    
    Audit   int64 `json:"audit,omitempty"`
    

    // 消息ID
    
    Mid   int64 `json:"mid,omitempty"`
    

    // 消息达到设备数
    
    Sentcount   int64 `json:"sentcount,omitempty"`
    

}
*/

// MessageResult 
type MessageResult struct {

    // 消息过期时间
    ExpireTime   int64 `json:"expire_time,omitempty"`

    // 消息预计发送数
    Total2send   int64 `json:"total2send,omitempty"`

    // 消息审核状态1,3:待审核 2:审核通过，处于发送中 4:审核拒绝 5 已撤回 6 已去重
    Audit   int64 `json:"audit,omitempty"`

    // 消息ID
    Mid   int64 `json:"mid,omitempty"`

    // 消息达到设备数
    Sentcount   int64 `json:"sentcount,omitempty"`

}
