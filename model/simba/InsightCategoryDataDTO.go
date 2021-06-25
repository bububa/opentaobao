package simba

// InsightCategoryDataDTO 
type InsightCategoryDataDTO struct {

    // 展现量
    Impression   int64 `json:"impression,omitempty"`

    // 点击量
    Click   int64 `json:"click,omitempty"`

    // 点击率
    Ctr   string `json:"ctr,omitempty"`

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 类目名称
    CatName   string `json:"cat_name,omitempty"`

}
