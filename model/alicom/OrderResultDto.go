package alicom

// OrderResultDto 
type OrderResultDto struct {

    // 结果描述
    Desc   string `json:"desc,omitempty"`

    // 天猫交易订单号
    OrderNo   int64 `json:"order_no,omitempty"`

    // 结果码
    ResultCode   string `json:"result_code,omitempty"`

    // 商家处理结果是否成功标志
    Success   bool `json:"success,omitempty"`

}
