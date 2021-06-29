package openmall

// TopItemVo 
type TopItemVo struct {

    // 商品类目
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 商品所在城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 商品成本价，精确到2位小数，单位：元。如：200.07，表示200元7分。
    
    CostPrice   string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
    

    // 商品描述, 字数要大于5个字符，小于25000个字符
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品图片
    
    ItemImages   []TopItemImageVo `json:"item_images,omitempty" xml:"item_images,omitempty"`
    

    // 商品视频列表，多个视频用逗号分隔。
    
    ItemVideos   string `json:"item_videos,omitempty" xml:"item_videos,omitempty"`
    

    // 商品主图
    
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    

    // 运费列表
    
    Postages   []TopPostageVo `json:"postages,omitempty" xml:"postages,omitempty"`
    

    // 商品的属性列表，由属性名ID(pid)、属性值ID(vid)、属性名(pid_name)、属性值(vid_name)组成。格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……
    
    ItemProperties   string `json:"item_properties,omitempty" xml:"item_properties,omitempty"`
    

    // 商品属性图片
    
    PropertyImages   []TopItemImageVo `json:"property_images,omitempty" xml:"property_images,omitempty"`
    

    // 商品所在省份
    
    Prov   string `json:"prov,omitempty" xml:"prov,omitempty"`
    

    // 库存
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 表示商品的体积，用于按体积计费的运费模板。该值的单位为立方米（m3）。该值支持两种格式的设置：  格式1：bulk:3，单位为立方米(m3)，表示直接设置为商品的体积；    格式2：width:10;breadth:10;height:10，单位为米（m）。
    
    Size   string `json:"size,omitempty" xml:"size,omitempty"`
    

    // sku列表
    
    Skus   []TopItemSkuVo `json:"skus,omitempty" xml:"skus,omitempty"`
    

    // 商品状态：上架（1）、下架（0）
    
    ItemStatus   int64 `json:"item_status,omitempty" xml:"item_status,omitempty"`
    

    // 商品标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 是否虚拟商品
    
    IsVirtual   bool `json:"is_virtual,omitempty" xml:"is_virtual,omitempty"`
    

    // 商品的重量，用于按重量计费的运费模板。注意：单位为kg
    
    Weight   string `json:"weight,omitempty" xml:"weight,omitempty"`
    

    // 是否区域限购。值为true时，通过taobao.openmall.item.salearea.get获取商品可销售区域；值为false时，该商品所有区域都可销售。
    
    AreaLimit   bool `json:"area_limit,omitempty" xml:"area_limit,omitempty"`
    

    // 属性值别名,比如颜色的自定义名称
    
    PropertyAlias   string `json:"property_alias,omitempty" xml:"property_alias,omitempty"`
    

    // 商品店铺名称
    
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    

    // 商品人气
    
    Popularity   int64 `json:"popularity,omitempty" xml:"popularity,omitempty"`
    

    // 当前商品是否支持hold单，当为false时，下单接口中的need_erp_hold不生效
    
    SupportErpHold   string `json:"support_erp_hold,omitempty" xml:"support_erp_hold,omitempty"`
    

    // 商品状态：上架（1）、下架（0）
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}
