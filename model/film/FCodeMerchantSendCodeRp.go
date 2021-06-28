package film

// FCodeMerchantSendCodeRp 
/* model for simplify = false
type FCodeMerchantSendCodeRp struct {

    // 码对外信息描述列表
    
    FCodeMerchantInfoList  struct {
        FCodeMerchantVo  []FCodeMerchantVo `json:"f_code_merchant_vo,omitempty"`
    } `json:"f_code_merchant_info_list,omitempty"`
    

}
*/

// FCodeMerchantSendCodeRp 
type FCodeMerchantSendCodeRp struct {

    // 码对外信息描述列表
    FCodeMerchantInfoList   []FCodeMerchantVo `json:"f_code_merchant_info_list,omitempty"`

}
