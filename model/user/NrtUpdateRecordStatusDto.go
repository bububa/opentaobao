package user

// NrtUpdateRecordStatusDto 
type NrtUpdateRecordStatusDto struct {

    // 客资记录ID
    RecordId   int64 `json:"record_id,omitempty"`

    // 创建人ID
    CreatorId   int64 `json:"creator_id,omitempty"`

    // 创建人姓名
    CreatorName   string `json:"creator_name,omitempty"`

    // 导购员ID
    EmployeeId   int64 `json:"employee_id,omitempty"`

    // 店铺ID
    StoreId   int64 `json:"store_id,omitempty"`

    // 2-已分配  3-已邀约  4-已到店  5-已下单  6-已付款
    Status   int64 `json:"status,omitempty"`

}
