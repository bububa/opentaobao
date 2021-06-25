package scbp

// ProductReportOperationDto 
type ProductReportOperationDto struct {

    // 产品名称或产品ID(模糊搜索)
    Key   string `json:"key,omitempty"`

    // 获取明细数据（"true"/"false"）,如果为"true"则为明细数据
    GetDetailData   string `json:"get_detail_data,omitempty"`

    // 开始时间(yyyy-MM-dd)
    DateBegin   string `json:"date_begin,omitempty"`

    // 结束时间(yyyy-MM-dd)
    DateEnd   string `json:"date_end,omitempty"`

    // 时间段
    DateRange   int64 `json:"date_range,omitempty"`

    // 营销目标
    CampaignType   int64 `json:"campaign_type,omitempty"`

    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 排序字段(impr/click/ctr/cost/cpc)
    OrderField   string `json:"order_field,omitempty"`

    // 排序方式(asc/desc)
    OrderType   string `json:"order_type,omitempty"`

    // 页码
    PageIndex   int64 `json:"page_index,omitempty"`

    // 每页数量
    PageSize   int64 `json:"page_size,omitempty"`

}
