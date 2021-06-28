package mos

// ResultDO 
/* model for simplify = false
type ResultDO struct {

    // data
    
    Data  *struct {
        CallDispatcherResponse  *CallDispatcherResponse `json:"call_dispatcher_response,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// ResultDO 
type ResultDO struct {

    // data
    Data   *CallDispatcherResponse `json:"data,omitempty"`

}
