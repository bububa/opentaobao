package baichuan

// TaobaoBaichuanItemSubscribeDailyLeftQueryResult 
/* model for simplify = false
type TaobaoBaichuanItemSubscribeDailyLeftQueryResult struct {

    // 返回
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemSubscribeDailyLeftQueryResult 
type TaobaoBaichuanItemSubscribeDailyLeftQueryResult struct {

    // 返回
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
