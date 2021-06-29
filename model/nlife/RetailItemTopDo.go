package nlife

// RetailItemTopDO 
type RetailItemTopDO struct {

    // 门店ID类型（零售+门店、商户中心门店、门店设备号
    
    StoreIdType   string `json:"store_id_type,omitempty" xml:"store_id_type,omitempty"`
    

    // 门店Id
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 类目id
    
    Cid   int64 `json:"cid,omitempty" xml:"cid,omitempty"`
    

    // 类目名称
    
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    

    // 品牌id
    
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    

    // 商品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 商家编码
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 条码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 货号/款号
    
    GoodsNo   string `json:"goods_no,omitempty" xml:"goods_no,omitempty"`
    

    // 商品库存
    
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    

    // 价格,单位元,保留2位小数
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 商品主图地址
    
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    

    // 商品图片列表,多个图片以逗号分隔
    
    Images   string `json:"images,omitempty" xml:"images,omitempty"`
    

    // 是否支持网直供
    
    SupportWzg   bool `json:"support_wzg,omitempty" xml:"support_wzg,omitempty"`
    

    // 是否支持淘宝客
    
    SupportTaoke   bool `json:"support_taoke,omitempty" xml:"support_taoke,omitempty"`
    

    // sku列表
    
    SkuList   []RetailSkuTopDO `json:"sku_list,omitempty" xml:"sku_list,omitempty"`
    

    // 商品的挂牌价-单位元,保留2位小数
    
    TagPrice   string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
    

    // 主图以外的图片列表
    
    ImageList   []string `json:"image_list,omitempty" xml:"image_list>string,omitempty"`
    

    // 无线端的商品详情
    
    WirelessDesc   string `json:"wireless_desc,omitempty" xml:"wireless_desc,omitempty"`
    

    // pcDesc
    
    PcDesc   string `json:"pc_desc,omitempty" xml:"pc_desc,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品的销售属性
    
    Props   string `json:"props,omitempty" xml:"props,omitempty"`
    

    // 详情页地址
    
    DetailUrl   string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
    

    // 商品的销售属性名称
    
    PropsName   string `json:"props_name,omitempty" xml:"props_name,omitempty"`
    

    // 商品类型: 0-IC线上商品; 1-商户导入线下商品
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // 类目树的信息(自顶向下),格式为 cid1:categoryName1;cid2:categoryName2;cid3:categoryName3
    
    AllCategoryInfo   string `json:"all_category_info,omitempty" xml:"all_category_info,omitempty"`
    

    // 网直供库存
    
    OnlineNum   int64 `json:"online_num,omitempty" xml:"online_num,omitempty"`
    

    // 供应商的ID
    
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    

    // 供应商的名称
    
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    

    // 零售商线上商品的itemId
    
    TaobaoItemId   int64 `json:"taobao_item_id,omitempty" xml:"taobao_item_id,omitempty"`
    

}
