package deliveryvoucher

// WatchAppointmentRequest 
type WatchAppointmentRequest struct {
    // 可扩展字段
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 流水号:唯一，幂等性判断
    OpId   string `json:"op_id,omitempty" xml:"op_id,omitempty"`
    // 操作时间
    OperationTime   string `json:"operation_time,omitempty" xml:"operation_time,omitempty"`
    // 第三方服务商标识 top appkey
    Provider   string `json:"provider,omitempty" xml:"provider,omitempty"`
    // 日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 数据类型0：初始化，1：修改，2：每天同步	必填
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 当前预约容量上限，type为0和1是必填
    CurrentCapacity   int64 `json:"current_capacity,omitempty" xml:"current_capacity,omitempty"`
    // 商家名称
    MerchantName   string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
    // 商家id
    MerchantId   int64 `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
    // 修改前预约容量上限	type为0和1是必填
    PreviousCapacity   int64 `json:"previous_capacity,omitempty" xml:"previous_capacity,omitempty"`
    // 最终预约数量	type为2必填
    ActualAppointment   int64 `json:"actual_appointment,omitempty" xml:"actual_appointment,omitempty"`
}
