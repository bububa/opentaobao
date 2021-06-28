package util

// ListDto 
/* model for simplify = false
type ListDto struct {

    // 申请单总数
    
    Count   int64 `json:"count,omitempty"`
    

    // 申请单列表
    
    List  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// ListDto 
type ListDto struct {

    // 申请单总数
    Count   int64 `json:"count,omitempty"`

    // 申请单列表
    List   []string `json:"list,omitempty"`

}
