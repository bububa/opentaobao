package tmallservice

// SettleAdjustmentResponse 
type SettleAdjustmentResponse struct {

    // cost，单位分
    Cost   int64 `json:"cost,omitempty"`

    // description
    Description   string `json:"description,omitempty"`

    // id
    Id   int64 `json:"id,omitempty"`

    // pictureUrls，多条已冒号分隔
    PictureUrls   string `json:"picture_urls,omitempty"`

    // priceFactors
    PriceFactors   string `json:"price_factors,omitempty"`

    // 调整单状态 待商家确认:1, 商家已确认:2,  待小二判定:3,  小二判定有效:4,  小二判定无效:5,  小二无法判定:6, 服务商取消:7, 超时确认:8, 完成:9
    Status   int64 `json:"status,omitempty"`

    // 工单ID
    WorkcardId   int64 `json:"workcard_id,omitempty"`

    // comments
    Comments   string `json:"comments,omitempty"`

    // gmtCreate
    CreateTime   string `json:"create_time,omitempty"`

    // gmtModified
    ModifiedTime   string `json:"modified_time,omitempty"`

    // serviceOrderId
    ServiceOrderId   int64 `json:"service_order_id,omitempty"`

    // 调整单类型
    Type   int64 `json:"type,omitempty"`

}
