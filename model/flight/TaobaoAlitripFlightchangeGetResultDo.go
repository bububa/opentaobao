package flight

// TaobaoAlitripFlightchangeGetResultDo 
/* model for simplify = false
type TaobaoAlitripFlightchangeGetResultDo struct {

    // results
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoAlitripFlightchangeGetResultDo 
type TaobaoAlitripFlightchangeGetResultDo struct {

    // results
    Results   []string `json:"results,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
