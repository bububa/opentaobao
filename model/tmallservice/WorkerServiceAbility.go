package tmallservice

// WorkerServiceAbility 
/* model for simplify = false
type WorkerServiceAbility struct {

    // 工人可服务区域
    
    AreaCodeList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"area_code_list,omitempty"`
    

}
*/

// WorkerServiceAbility 
type WorkerServiceAbility struct {

    // 工人可服务区域
    AreaCodeList   []int64 `json:"area_code_list,omitempty"`

}
