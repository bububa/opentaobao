package drugtrace

// CodeActiveStatusDTO 
type CodeActiveStatusDTO struct {
    // 最后激活时间，到毫秒时间timeMills方式
    ActiveTime   int64 `json:"active_time,omitempty" xml:"active_time,omitempty"`
    // 码段
    ResProdCode   string `json:"res_prod_code,omitempty" xml:"res_prod_code,omitempty"`
}
