package promotion

// AlibabaBenefitQueryResult 
type AlibabaBenefitQueryResult struct {

    // msg
    Msg   string `json:"msg,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // datas
    Datas   []OrightDto `json:"datas,omitempty"`

}
