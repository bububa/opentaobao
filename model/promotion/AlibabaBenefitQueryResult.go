package promotion

// AlibabaBenefitQueryResult 
/* model for simplify = false
type AlibabaBenefitQueryResult struct {

    // msg
    
    Msg   string `json:"msg,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // datas
    
    Datas  struct {
        OrightDto  []OrightDto `json:"oright_dto,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

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
