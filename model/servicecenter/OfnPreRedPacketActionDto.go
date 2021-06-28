package servicecenter

// OfnPreRedPacketActionDto 
/* model for simplify = false
type OfnPreRedPacketActionDto struct {

    // 主键
    
    Id   int64 `json:"id,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 操作类型。1=天猫发预付红包；2=天猫发尾款红包；3=天猫扣回红包；4=回收商扣回红包
    
    ActionType   int64 `json:"action_type,omitempty"`
    

    // 状态。初始化=1，重试中=2，失败=3，成功=4
    
    Status   int64 `json:"status,omitempty"`
    

    // 资金池的记录
    
    AfterFundRecordList  struct {
        OfnPreRedPacketFundRecordDto  []OfnPreRedPacketFundRecordDto `json:"ofn_pre_red_packet_fund_record_dto,omitempty"`
    } `json:"after_fund_record_list,omitempty"`
    

}
*/

// OfnPreRedPacketActionDto 
type OfnPreRedPacketActionDto struct {

    // 主键
    Id   int64 `json:"id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 操作类型。1=天猫发预付红包；2=天猫发尾款红包；3=天猫扣回红包；4=回收商扣回红包
    ActionType   int64 `json:"action_type,omitempty"`

    // 状态。初始化=1，重试中=2，失败=3，成功=4
    Status   int64 `json:"status,omitempty"`

    // 资金池的记录
    AfterFundRecordList   []OfnPreRedPacketFundRecordDto `json:"after_fund_record_list,omitempty"`

}
