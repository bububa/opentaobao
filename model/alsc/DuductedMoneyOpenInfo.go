package alsc

// DuductedMoneyOpenInfo 
/* model for simplify = false
type DuductedMoneyOpenInfo struct {

    // 消耗的积分
    
    ConsumePoint   int64 `json:"consume_point,omitempty"`
    

    // 抵扣的金额,单位为分
    
    DeductedMoney   int64 `json:"deducted_money,omitempty"`
    

    // 是否删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

}
*/

// DuductedMoneyOpenInfo 
type DuductedMoneyOpenInfo struct {

    // 消耗的积分
    ConsumePoint   int64 `json:"consume_point,omitempty"`

    // 抵扣的金额,单位为分
    DeductedMoney   int64 `json:"deducted_money,omitempty"`

    // 是否删除
    Deleted   bool `json:"deleted,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
