package tvupadmin

// UserRightsDO 
type UserRightsDO struct {
    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 用户ID
    Uid   int64 `json:"uid,omitempty" xml:"uid,omitempty"`
    // 是否连续包月
    RenewalSupport   bool `json:"renewal_support,omitempty" xml:"renewal_support,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 权益截止时间
    GmtEnd   string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
    // 权益开始时间
    GmtStart   string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
    // 权益类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 权益ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
