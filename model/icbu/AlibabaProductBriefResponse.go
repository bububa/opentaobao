package icbu

// AlibabaProductBriefResponse 
type AlibabaProductBriefResponse struct {

    // 商品ID
    Id   int64 `json:"id,omitempty"`

    // 混淆后的商品ID
    ProductId   string `json:"product_id,omitempty"`

    // 分组ID
    GroupId   int64 `json:"group_id,omitempty"`

    // 分组名称
    GroupName   string `json:"group_name,omitempty"`

    // 商品名称
    Subject   string `json:"subject,omitempty"`

    // 商品状态
    Status   string `json:"status,omitempty"`

    // 关键词
    Keywords   []String `json:"keywords,omitempty"`

    // 商品的主图
    MainImage   *MainImage `json:"main_image,omitempty"`

    // 商品类型
    ProductType   string `json:"product_type,omitempty"`

    // 语种
    Language   string `json:"language,omitempty"`

    // Y为上架状态
    Display   string `json:"display,omitempty"`

    // 最近一次修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 类目ID
    CategoryId   int64 `json:"category_id,omitempty"`

    // james
    OwnerMemberDisplayName   string `json:"owner_member_display_name,omitempty"`

    // true
    IsSpecific   bool `json:"is_specific,omitempty"`

    // true
    IsRts   bool `json:"is_rts,omitempty"`

    // https://www.alibaba.com/product-detail/Eco-Friendly-100-Biodegradable-Cornstarch-Trash_60832548452.html?spm=a2700.galleryofferlist.normalList.12.6c612db4ueHAW2&fullFirstScreen=true
    PcDetailUrl   string `json:"pc_detail_url,omitempty"`

    // true
    SmartEdit   bool `json:"smart_edit,omitempty"`

    // 2020-12-22 12:00:00
    GmtCreate   string `json:"gmt_create,omitempty"`

}
