package drugtrace

// AlibabaAlihealthDrugKytDrBillcheckModel 
type AlibabaAlihealthDrugKytDrBillcheckModel struct {

    // 单据日期
    
    BillTime   string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
    

    // 单据类型
    
    BillType   int64 `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    

    // 追溯验证数量
    
    MatchedCount   int64 `json:"matched_count,omitempty" xml:"matched_count,omitempty"`
    

    // 收发货企业名称
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 验证率
    
    MatchedRateShow   string `json:"matched_rate_show,omitempty" xml:"matched_rate_show,omitempty"`
    

    // 单据编号
    
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    

    // 码验证信息
    
    DetailInfoList   []DetailInfoList `json:"detail_info_list,omitempty" xml:"detail_info_list,omitempty"`
    

}
