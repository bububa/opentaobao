package flight

// TaobaoAlitripFlightchangeGetResultDo 
type TaobaoAlitripFlightchangeGetResultDo struct {

    // results
    Results   []Json `json:"results,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
