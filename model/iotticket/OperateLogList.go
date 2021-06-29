package iotticket

// OperateLogList 
type OperateLogList struct {
    // 操作人联系方式
    OperatePhone   string `json:"operate_phone,omitempty" xml:"operate_phone,omitempty"`
    // 操作时间
    OperateTime   string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
    // 操作类型描述
    OperateType   string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
    // 操作人
    OperateName   string `json:"operate_name,omitempty" xml:"operate_name,omitempty"`
    // 备注信息
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
}
