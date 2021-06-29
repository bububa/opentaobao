package wdk

// OrderBalanceBillRequest 
type OrderBalanceBillRequest struct {
    // 业务日期
    Thedate   string `json:"thedate,omitempty" xml:"thedate,omitempty"`
    // 每页条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // storeId
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 最大的id
    MaxId   int64 `json:"max_id,omitempty" xml:"max_id,omitempty"`
}
