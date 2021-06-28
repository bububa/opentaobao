package tanx

// CreativePublisherAuditDto 
/* model for simplify = false
type CreativePublisherAuditDto struct {

    // 状态（PASS.通过；REFUSE.拒绝；WAITING.待审）
    
    Status   string `json:"status,omitempty"`
    

    // 媒体ID
    
    PubliserId   string `json:"publiser_id,omitempty"`
    

    // 拒绝原因
    
    RefuseCause   string `json:"refuse_cause,omitempty"`
    

}
*/

// CreativePublisherAuditDto 
type CreativePublisherAuditDto struct {

    // 状态（PASS.通过；REFUSE.拒绝；WAITING.待审）
    Status   string `json:"status,omitempty"`

    // 媒体ID
    PubliserId   string `json:"publiser_id,omitempty"`

    // 拒绝原因
    RefuseCause   string `json:"refuse_cause,omitempty"`

}
