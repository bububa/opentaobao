package mos

// RightsList 
/* model for simplify = false
type RightsList struct {

    // 券价值
    
    CouponAmount   int64 `json:"coupon_amount,omitempty"`
    

    // id
    
    SnapshotId   int64 `json:"snapshot_id,omitempty"`
    

    // 券使用门槛
    
    EntryAmount   int64 `json:"entry_amount,omitempty"`
    

    // 券名称
    
    Name   string `json:"name,omitempty"`
    

    // 开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

}
*/

// RightsList 
type RightsList struct {

    // 券价值
    CouponAmount   int64 `json:"coupon_amount,omitempty"`

    // id
    SnapshotId   int64 `json:"snapshot_id,omitempty"`

    // 券使用门槛
    EntryAmount   int64 `json:"entry_amount,omitempty"`

    // 券名称
    Name   string `json:"name,omitempty"`

    // 开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 结束时间
    EndTime   string `json:"end_time,omitempty"`

}
