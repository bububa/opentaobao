package promotion

// CouponTemplateQueryRequest 
type CouponTemplateQueryRequest struct {

    // 模板表ID
    Id   int64 `json:"id,omitempty"`

    // ump模板ID
    SourceId   int64 `json:"source_id,omitempty"`

    // 用户信息
    UserInfo   *UserInfo `json:"user_info,omitempty"`

}
