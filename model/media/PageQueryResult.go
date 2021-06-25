package media

// PageQueryResult 
type PageQueryResult struct {

    // 总数
    Total   int64 `json:"total,omitempty"`

    // resultList
    ResultList   []Resultlist `json:"result_list,omitempty"`

}
