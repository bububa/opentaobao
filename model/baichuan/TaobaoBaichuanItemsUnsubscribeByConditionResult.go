package baichuan

// TaobaoBaichuanItemsUnsubscribeByConditionResult 
/* model for simplify = false
type TaobaoBaichuanItemsUnsubscribeByConditionResult struct {

    // 分返回码返回结果
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemsUnsubscribeByConditionResult 
type TaobaoBaichuanItemsUnsubscribeByConditionResult struct {

    // 分返回码返回结果
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
