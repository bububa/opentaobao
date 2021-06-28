package promotion

// ActivityRelationDetailRequest 
/* model for simplify = false
type ActivityRelationDetailRequest struct {

    // 活动状态(VALID  ， DELETE)
    
    Status   string `json:"status,omitempty"`
    

    // ISV活动关联权益后获得的关联ID
    
    RelationId   int64 `json:"relation_id,omitempty"`
    

}
*/

// ActivityRelationDetailRequest 
type ActivityRelationDetailRequest struct {

    // 活动状态(VALID  ， DELETE)
    Status   string `json:"status,omitempty"`

    // ISV活动关联权益后获得的关联ID
    RelationId   int64 `json:"relation_id,omitempty"`

}
