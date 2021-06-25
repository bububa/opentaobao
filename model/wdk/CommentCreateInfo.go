package wdk

// CommentCreateInfo 
type CommentCreateInfo struct {

    // 商家评价Id
    OutCommentId   string `json:"out_comment_id,omitempty"`

    // 履约门店名称
    StoreName   string `json:"store_name,omitempty"`

    // 履约门店Id
    StoreId   string `json:"store_id,omitempty"`

    // 盒马主单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 盒马子单号
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty"`

    // 类目名称
    Category   string `json:"category,omitempty"`

    // 商品金额
    Price   int64 `json:"price,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 评价日期
    CommentDate   string `json:"comment_date,omitempty"`

    // 评价星级
    CommentStar   string `json:"comment_star,omitempty"`

    // 评价内容
    CommentContent   string `json:"comment_content,omitempty"`

    // 评价图片
    CommentPics   string `json:"comment_pics,omitempty"`

    // 评价标签
    CommentTag   string `json:"comment_tag,omitempty"`

    // 差评原因
    CommentReason   string `json:"comment_reason,omitempty"`

    // 差评回评
    ReviewReason   string `json:"review_reason,omitempty"`

}
