package baichuan

// TaobaoBaichuanItemSubscribeRelationQueryResult 
/* model for simplify = false
type TaobaoBaichuanItemSubscribeRelationQueryResult struct {

    // 返回的list
    
    ResultList  struct {
        ResultMeta  []ResultMeta `json:"result_meta,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoBaichuanItemSubscribeRelationQueryResult 
type TaobaoBaichuanItemSubscribeRelationQueryResult struct {

    // 返回的list
    ResultList   []ResultMeta `json:"result_list,omitempty"`

}
