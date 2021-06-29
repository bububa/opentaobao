package omniorder

// TaobaoOmniorderStorecollectQueryResult 
type TaobaoOmniorderStorecollectQueryResult struct {

    // 0表示码可用，其余值表示码不可用
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // data
    
    Data   *StoreCollectQueryOrderResponse `json:"data,omitempty" xml:"data,omitempty"`
    

    // 码状态附加信息，若码可用则此处为空
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
