package media

// PageQueryResult 
/* model for simplify = false
type PageQueryResult struct {

    // 总数
    
    Total   int64 `json:"total,omitempty"`
    

    // resultList
    
    ResultList  struct {
        Resultlist  []Resultlist `json:"resultlist,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// PageQueryResult 
type PageQueryResult struct {

    // 总数
    Total   int64 `json:"total,omitempty"`

    // resultList
    ResultList   []Resultlist `json:"result_list,omitempty"`

}
