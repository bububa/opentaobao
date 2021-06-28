package baichuan

// TaobaoBaichuanItemsSubscribeResult 
/* model for simplify = false
type TaobaoBaichuanItemsSubscribeResult struct {

    // 按不同的返回码将结果分部分返回
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemsSubscribeResult 
type TaobaoBaichuanItemsSubscribeResult struct {

    // 按不同的返回码将结果分部分返回
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
