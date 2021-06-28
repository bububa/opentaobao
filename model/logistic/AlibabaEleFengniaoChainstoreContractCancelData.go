package logistic

// AlibabaEleFengniaoChainstoreContractCancelData 
/* model for simplify = false
type AlibabaEleFengniaoChainstoreContractCancelData struct {

    // appid
    
    AppId   string `json:"app_id,omitempty"`
    

    // 商户code
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 门店code
    
    ChainstoreCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"chainstore_codes,omitempty"`
    

}
*/

// AlibabaEleFengniaoChainstoreContractCancelData 
type AlibabaEleFengniaoChainstoreContractCancelData struct {

    // appid
    AppId   string `json:"app_id,omitempty"`

    // 商户code
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 门店code
    ChainstoreCodes   []string `json:"chainstore_codes,omitempty"`

}
