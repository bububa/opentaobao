package wdk

// ReturnBoxContainerRequest 
/* model for simplify = false
type ReturnBoxContainerRequest struct {

    // 周转箱列表
    
    BoxCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"box_code_list,omitempty"`
    

    // 收箱交接单号
    
    HandOverNo   string `json:"hand_over_no,omitempty"`
    

    // 仓编号
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 操作时间
    
    OperateTime   string `json:"operate_time,omitempty"`
    

}
*/

// ReturnBoxContainerRequest 
type ReturnBoxContainerRequest struct {

    // 周转箱列表
    BoxCodeList   []string `json:"box_code_list,omitempty"`

    // 收箱交接单号
    HandOverNo   string `json:"hand_over_no,omitempty"`

    // 仓编号
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 操作时间
    OperateTime   string `json:"operate_time,omitempty"`

}
