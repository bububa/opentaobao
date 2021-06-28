package scbp

// AccountReportDto 
/* model for simplify = false
type AccountReportDto struct {

    // 返回数据集合
    
    AccountEffectList  struct {
        AccountEffectDto  []AccountEffectDto `json:"account_effect_dto,omitempty"`
    } `json:"account_effect_list,omitempty"`
    

}
*/

// AccountReportDto 
type AccountReportDto struct {

    // 返回数据集合
    AccountEffectList   []AccountEffectDto `json:"account_effect_list,omitempty"`

}
