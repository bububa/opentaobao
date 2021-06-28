package tmallservice

// ServiceMonitorMessage 
/* model for simplify = false
type ServiceMonitorMessage struct {

    // 消息id
    
    Id   int64 `json:"id,omitempty"`
    

    // 提醒文本
    
    Content   string `json:"content,omitempty"`
    

    // 预警消息级别 1、预警 2、警告 3、严重
    
    Level   int64 `json:"level,omitempty"`
    

    // 备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 消息创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 状态 0、已生成 1、已预警 2、已收到 3、已读
    
    Status   int64 `json:"status,omitempty"`
    

    // 规则类型
    
    RuleId   string `json:"rule_id,omitempty"`
    

    // 服务类型
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // bizId的业务实体类型，如果为1，则bizId为工单id
    
    BizType   int64 `json:"biz_type,omitempty"`
    

    // 业务实体id
    
    BizId   int64 `json:"biz_id,omitempty"`
    

}
*/

// ServiceMonitorMessage 
type ServiceMonitorMessage struct {

    // 消息id
    Id   int64 `json:"id,omitempty"`

    // 提醒文本
    Content   string `json:"content,omitempty"`

    // 预警消息级别 1、预警 2、警告 3、严重
    Level   int64 `json:"level,omitempty"`

    // 备注
    Memo   string `json:"memo,omitempty"`

    // 消息创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 状态 0、已生成 1、已预警 2、已收到 3、已读
    Status   int64 `json:"status,omitempty"`

    // 规则类型
    RuleId   string `json:"rule_id,omitempty"`

    // 服务类型
    ServiceCode   string `json:"service_code,omitempty"`

    // bizId的业务实体类型，如果为1，则bizId为工单id
    BizType   int64 `json:"biz_type,omitempty"`

    // 业务实体id
    BizId   int64 `json:"biz_id,omitempty"`

}
