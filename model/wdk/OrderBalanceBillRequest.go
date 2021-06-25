package wdk

// OrderBalanceBillRequest 
type OrderBalanceBillRequest struct {

    // 业务日期
    Thedate   string `json:"thedate,omitempty"`

    // 每页条数
    PageSize   int64 `json:"page_size,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // storeId
    StoreId   string `json:"store_id,omitempty"`

    // 最大的id
    MaxId   int64 `json:"max_id,omitempty"`

}
