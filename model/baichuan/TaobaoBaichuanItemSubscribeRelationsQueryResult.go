package baichuan

// TaobaoBaichuanItemSubscribeRelationsQueryResult 
/* model for simplify = false
type TaobaoBaichuanItemSubscribeRelationsQueryResult struct {

    // 返回的list
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemSubscribeRelationsQueryResult 
type TaobaoBaichuanItemSubscribeRelationsQueryResult struct {

    // 返回的list
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
