package wdk

// CashierShiftFundRequest 
/* model for simplify = false
type CashierShiftFundRequest struct {

    // 门店编号 示例 DRF4012（优先使用）
    
    ShopCode   string `json:"shop_code,omitempty"`
    

    // 门店Id（可选，shopCode为空时使用shopId）
    
    ShopId   int64 `json:"shop_id,omitempty"`
    

    // 业务日期 yyyy-MM-dd 格式（优先使用）
    
    BizDate   string `json:"biz_date,omitempty"`
    

    // 开始时间 yyyy-MM-dd HH:mm:ss(可选，bizDate为空时使用startTime)
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 结束时间 yyyy-MM-dd HH:mm:ss(可选，bizDate为空时使用endTime)
    
    EndTime   string `json:"end_time,omitempty"`
    

}
*/

// CashierShiftFundRequest 
type CashierShiftFundRequest struct {

    // 门店编号 示例 DRF4012（优先使用）
    ShopCode   string `json:"shop_code,omitempty"`

    // 门店Id（可选，shopCode为空时使用shopId）
    ShopId   int64 `json:"shop_id,omitempty"`

    // 业务日期 yyyy-MM-dd 格式（优先使用）
    BizDate   string `json:"biz_date,omitempty"`

    // 开始时间 yyyy-MM-dd HH:mm:ss(可选，bizDate为空时使用startTime)
    StartTime   string `json:"start_time,omitempty"`

    // 结束时间 yyyy-MM-dd HH:mm:ss(可选，bizDate为空时使用endTime)
    EndTime   string `json:"end_time,omitempty"`

}
