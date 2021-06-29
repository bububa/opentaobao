package film

// FCodeMerchantVo 
type FCodeMerchantVo struct {
    // 码过期时间
    GmtExpire   string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
    // 码生成任务ID
    GenTaskId   int64 `json:"gen_task_id,omitempty" xml:"gen_task_id,omitempty"`
    // 码可抵用金额
    CostPrice   int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
