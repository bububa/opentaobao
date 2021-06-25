package mos

// ParameterEntity 
type ParameterEntity struct {

    // 流程定义KEY,发起流程的唯一标
    ProcessDefinitionKey   string `json:"process_definition_key,omitempty"`

    // 流程审批类型
    ProcessType   string `json:"process_type,omitempty"`

    // 业务详情页URL
    BusinessDataUrl   string `json:"business_data_url,omitempty"`

    // 发起人角色名称
    StartUserRoleName   string `json:"start_user_role_name,omitempty"`

    // 操作人ID
    OperId   string `json:"oper_id,omitempty"`

    // 业务表单ID(系统-业务类型-编号)
    FormId   string `json:"form_id,omitempty"`

    // 标题
    Title   string `json:"title,omitempty"`

    // 门店NO
    StoreNo   string `json:"store_no,omitempty"`

    // 抄送人(非必传)
    CcUserIds   string `json:"cc_user_ids,omitempty"`

    // 发起人ID类型(1:淘宝UserId,2:MIS系统 OuterId)
    OperIdType   string `json:"oper_id_type,omitempty"`

    // 业种code
    YzCode   string `json:"yz_code,omitempty"`

    // 流程发起描述信息(会展示在审批日志中)
    Message   string `json:"message,omitempty"`

}
