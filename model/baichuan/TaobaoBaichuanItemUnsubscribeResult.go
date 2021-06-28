package baichuan

// TaobaoBaichuanItemUnsubscribeResult 
/* model for simplify = false
type TaobaoBaichuanItemUnsubscribeResult struct {

    // 返回按resultCode分为多个返回部分
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemUnsubscribeResult 
type TaobaoBaichuanItemUnsubscribeResult struct {

    // 返回按resultCode分为多个返回部分
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
