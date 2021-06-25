package product

// GoodsDetail 
type GoodsDetail struct {

    // max_price_cent产品最高价格
    MaxPriceCent   int64 `json:"max_price_cent,omitempty"`

    // detail_url 产品detail页面地址
    DetailUrl   string `json:"detail_url,omitempty"`

    // weight 产品重量
    Weight   string `json:"weight,omitempty"`

    // subject 产品名
    Subject   string `json:"subject,omitempty"`

    // keyword 产品关键词
    Keyword   string `json:"keyword,omitempty"`

    // image_urls 产品图片地址
    ImageUrls   []String `json:"image_urls,omitempty"`

    // thumb_urls 产品缩略图地址
    ThumbUrls   []String `json:"thumb_urls,omitempty"`

    // property_list 产品属性
    PropertyList   []WholesaleGoodsAttribute `json:"property_list,omitempty"`

    // moq 产品最小起订量
    Moq   int64 `json:"moq,omitempty"`

    // inbox_num 产品入箱数
    InboxNum   int64 `json:"inbox_num,omitempty"`

    // package_method 产品包装方式
    PackageMethod   string `json:"package_method,omitempty"`

    // package_size 产品大小
    PackageSize   string `json:"package_size,omitempty"`

    // id 产品id
    Id   string `json:"id,omitempty"`

    // unit 产品最小起订量单位
    Unit   string `json:"unit,omitempty"`

    // description 产品详细描述
    Description   string `json:"description,omitempty"`

    // supplier_info
    SupplierInfo   *WholesaleUser `json:"supplier_info,omitempty"`

    // 产品物流信息
    FreightInfoList   []FreightInfo `json:"freight_info_list,omitempty"`

    // buy_now_url 产品下单链接
    BuyNowUrl   string `json:"buy_now_url,omitempty"`

    // batch_sale 产品售卖方式，如是否批量售卖
    BatchSale   bool `json:"batch_sale,omitempty"`

    // shipping_date  发货期
    ShippingDate   int64 `json:"shipping_date,omitempty"`

    // price_currency 产品价格单位
    Currency   string `json:"currency,omitempty"`

    // min_price_cent 产品最低价格
    MinPriceCent   int64 `json:"min_price_cent,omitempty"`

}
