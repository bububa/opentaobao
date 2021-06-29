package scbp

// TopP4pQuickCampaignQueryDTO 
type TopP4pQuickCampaignQueryDTO struct {
    // 第几页
    ToPage   int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
    // 每页返回数量
    PerPageSize   int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
}
