package baichuan

// TaobaoBaichuanItemSubscribeResult 
/* model for simplify = false
type TaobaoBaichuanItemSubscribeResult struct {

    // 返回的列表
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemSubscribeResult 
type TaobaoBaichuanItemSubscribeResult struct {

    // 返回的列表
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
