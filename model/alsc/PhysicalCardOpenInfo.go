package alsc

// PhysicalCardOpenInfo 
type PhysicalCardOpenInfo struct {

    // 绑定操作人ID
    
    BindOperatorId   string `json:"bind_operator_id,omitempty" xml:"bind_operator_id,omitempty"`
    

    // 绑定门店ID
    
    BindShopId   string `json:"bind_shop_id,omitempty" xml:"bind_shop_id,omitempty"`
    

    // 绑定时间
    
    BindTime   string `json:"bind_time,omitempty" xml:"bind_time,omitempty"`
    

    // 卡实例ID
    
    CardId   string `json:"card_id,omitempty" xml:"card_id,omitempty"`
    

    // 卡模板ID
    
    CardTemplateId   string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
    

    // 卡类型
    
    CardType   string `json:"card_type,omitempty" xml:"card_type,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty" xml:"create_by,omitempty"`
    

    // 逻辑删除
    
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 操作员id
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 物理卡号
    
    PhysicalCardId   string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
    

    // 会员计划id
    
    PlanId   string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
    

    // 发行记录ID
    
    PublishId   string `json:"publish_id,omitempty" xml:"publish_id,omitempty"`
    

    // WAIT_MAKE("WAIT_MAKE", "未制卡"),      /**      * 未出售      */     WAIT_SELL("WAIT_SELL", "未出售"),      /**      * 已出售      */     SOLD("SOLD", "已出售"),      /**      * 已作废      */     INVALID("INVALID", "已作废");
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 修改人
    
    UpdateBy   string `json:"update_by,omitempty" xml:"update_by,omitempty"`
    

    // 扩展信息
    
    ExtInfo   *Extinfo `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
    

}
