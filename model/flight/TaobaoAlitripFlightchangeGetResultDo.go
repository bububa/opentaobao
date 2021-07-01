package flight

// TaobaoAlitripFlightchangeGetResultDo 
type TaobaoAlitripFlightchangeGetResultDo struct {
    // results
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    // errMsg
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // errCode
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
