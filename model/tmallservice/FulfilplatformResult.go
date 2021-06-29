package tmallservice

// FulfilplatformResult 
type FulfilplatformResult struct {
    // 核销单id
    ResultData   int64 `json:"result_data,omitempty" xml:"result_data,omitempty"`
    // 是否执行成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 错误编码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 工单Id
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    // 展示信息
    DisplayMsg   string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
    // 错误名称
    ErrorName   string `json:"error_name,omitempty" xml:"error_name,omitempty"`
    // 服务单列表数据
    ResultList   []SpServiceOrderDTO `json:"result_list,omitempty" xml:"result_list>sp_service_order_dto,omitempty"`
    // 错误类型
    ErrorType   string `json:"error_type,omitempty" xml:"error_type,omitempty"`
    // 物流类型 消费者--> 商家 STAGE1；商家--->消费者  STAGE2
    StageType   string `json:"stage_type,omitempty" xml:"stage_type,omitempty"`
}
