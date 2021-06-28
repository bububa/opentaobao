package wdk

// UtmsResult 
/* model for simplify = false
type UtmsResult struct {

    // code
    
    Code   string `json:"code,omitempty"`
    

    // model
    
    Model   bool `json:"model,omitempty"`
    

    // msg
    
    Msg   string `json:"msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // list
    
    List  struct {
        BomProcessDTO  []BomProcessDTO `json:"bom_process_dto,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// UtmsResult 
type UtmsResult struct {

    // code
    Code   string `json:"code,omitempty"`

    // model
    Model   bool `json:"model,omitempty"`

    // msg
    Msg   string `json:"msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // list
    List   []BomProcessDTO `json:"list,omitempty"`

}
