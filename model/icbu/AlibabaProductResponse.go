package icbu

// AlibabaProductResponse 
type AlibabaProductResponse struct {

    // 类目ID
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 商品分组ID
    
    GroupId   int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
    

    // 商品名称
    
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    

    // status 的值：sketch：草稿，approved：审核通过，tbd：审核不通过，new 、modified ：审核中
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 关键词
    
    Keywords   []string `json:"keywords,omitempty" xml:"keywords>string,omitempty"`
    

    // 商品属性
    
    Attributes   []ProductAttribute `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 商品的主图
    
    MainImage   *MainImage `json:"main_image,omitempty" xml:"main_image,omitempty"`
    

    // 询盘商品交易信息
    
    SourcingTrade   *SourcingTrade `json:"sourcing_trade,omitempty" xml:"sourcing_trade,omitempty"`
    

    // 在线批发商品交易信息
    
    WholesaleTrade   *WholesaleTrade `json:"wholesale_trade,omitempty" xml:"wholesale_trade,omitempty"`
    

    // 商品SKU
    
    ProductSku   *ProductSkuResponse `json:"product_sku,omitempty" xml:"product_sku,omitempty"`
    

    // 商品详情描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 商品类型
    
    ProductType   string `json:"product_type,omitempty" xml:"product_type,omitempty"`
    

    // 语种
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // 产品ID
    
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 产品负责人
    
    OwnerMember   int64 `json:"owner_member,omitempty" xml:"owner_member,omitempty"`
    

    // 产品更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // Y为上架状态
    
    Display   string `json:"display,omitempty" xml:"display,omitempty"`
    

    // 产品负责人显示名，由firstname和lastname拼接组成
    
    OwnerMemberDisplayName   string `json:"owner_member_display_name,omitempty" xml:"owner_member_display_name,omitempty"`
    

    // 定制信息
    
    CustomInfo   *CustomInfo `json:"custom_info,omitempty" xml:"custom_info,omitempty"`
    

    // 是否是智能编辑
    
    IsSmartEdit   bool `json:"is_smart_edit,omitempty" xml:"is_smart_edit,omitempty"`
    

    // /**      * SKU价      */     SKU_PRICE("sku_price"),     /**      * 阶梯价      */     LADDER_PRICE("ladder_price"),     /**      * fob价: 单一区间fob价      */     FOB_PRICE("fob_price");
    
    PriceType   string `json:"price_type,omitempty" xml:"price_type,omitempty"`
    

    // https://www.alibaba.com/product-detail/Short-Umbrella-Girls-Black-Lace-Polka_1600107214049.html?spm=a2700.galleryofferlist.normalList.12.6c612db4ueHAW2
    
    PcDetailUrl   string `json:"pc_detail_url,omitempty" xml:"pc_detail_url,omitempty"`
    

}
