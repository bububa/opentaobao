package alsc

// PrizeItemInfo 
type PrizeItemInfo struct {

    // 优惠金额
    
    Denomination   string `json:"denomination,omitempty" xml:"denomination,omitempty"`
    

    // 几等奖
    
    PrizeLevel   int64 `json:"prize_level,omitempty" xml:"prize_level,omitempty"`
    

    // 奖品名称
    
    PrizeName   string `json:"prize_name,omitempty" xml:"prize_name,omitempty"`
    

    // 券id
    
    VoucherIds   []string `json:"voucher_ids,omitempty" xml:"voucher_ids>string,omitempty"`
    

    // 券名称
    
    VoucherName   string `json:"voucher_name,omitempty" xml:"voucher_name,omitempty"`
    

    // 券类型
    
    VoucherType   string `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
    

}
