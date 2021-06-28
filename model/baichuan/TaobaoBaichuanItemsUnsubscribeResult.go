package baichuan

// TaobaoBaichuanItemsUnsubscribeResult 
/* model for simplify = false
type TaobaoBaichuanItemsUnsubscribeResult struct {

    // 返回按resultCode分为多个返回部分
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemsUnsubscribeResult 
type TaobaoBaichuanItemsUnsubscribeResult struct {

    // 返回按resultCode分为多个返回部分
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
