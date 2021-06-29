package wdk

// UtmsResult 
type UtmsResult struct {
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // model
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // list
    List   []BomProcessDTO `json:"list,omitempty" xml:"list>bom_process_dto,omitempty"`
}
