package idleisv

// IdleItemApiDo 
type IdleItemApiDo struct {

    // 商品描述(长度<=5000)
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 商品Id（根据此数据进行相应商品更新）
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品原价, 单位：元(最大99999999)
    
    OriginalPrice   string `json:"original_price,omitempty" xml:"original_price,omitempty"`
    

    // 商品售价, 单位：元(最大99999999)
    
    ReservePrice   string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
    

    // 商品标题(长度<=20)
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 邮费, 单位：元(最大99999999)
    
    TransportFee   string `json:"transport_fee,omitempty" xml:"transport_fee,omitempty"`
    

    // 属性的键值对信息，包括品牌、型号、内存大小（手机）等，(不传入则不修改)
    
    PvList   []IdleNewPubValueDo `json:"pv_list,omitempty" xml:"pv_list,omitempty"`
    

    // sku列表(不传入则不修改)
    
    ItemSkuList   []IdleItemApiSkuDo `json:"item_sku_list,omitempty" xml:"item_sku_list,omitempty"`
    

    // 图书业务数据(不传入则不修改)
    
    BookData   *IdleItemApiBookDo `json:"book_data,omitempty" xml:"book_data,omitempty"`
    

    // 商品图片列表，使用图片上传接口返回的文件Id，支持多张(最多9张)
    
    Images   []int64 `json:"images,omitempty" xml:"images>int64,omitempty"`
    

    // 商品新旧程度, 值为0-100的整数，例如100代表全新，95代表95新；特殊场景会影响优品标
    
    StuffStatus   int64 `json:"stuff_status,omitempty" xml:"stuff_status,omitempty"`
    

    // 行政区划Id，最小行政单位code，若是地区级别，则为地区级别的id；否则为城市级别的id(6位)
    
    DivisionId   int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
    

    // 纬度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 经度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

    // 已验货业务数据
    
    InspectedData   *IdleItemApiInspectedDo `json:"inspected_data,omitempty" xml:"inspected_data,omitempty"`
    

    // 宝贝库存数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 类目Id
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 图片链接
    
    ImgUrls   []string `json:"img_urls,omitempty" xml:"img_urls>string,omitempty"`
    

    // 验货报告url链接(长度<=300)
    
    InspectReport   string `json:"inspect_report,omitempty" xml:"inspect_report,omitempty"`
    

    // 服务商商品业务分类，手机:1, 潮品:2, 家电:3, 乐器:8, 3C数码:9, 奢品:16, 母婴:17, 美妆:18, 文玩/珠宝:19, 潮玩:20, 家居:21
    
    SpBizType   string `json:"sp_biz_type,omitempty" xml:"sp_biz_type,omitempty"`
    

    // 商品状态，可选值为0（表示在线），1（表示售出下架）-2（表示下架），-1（表示删除），99（其他）
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 交易方式, 仅在线交易: 0,仅线下交易: 1,线上OR线下交易: 2（int型，1位）
    
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    

    // 入仓城市，不可修改
    
    WareHouseCity   string `json:"ware_house_city,omitempty" xml:"ware_house_city,omitempty"`
    

    // 入仓时间，时间戳秒，不可修改
    
    WareHouseTime   int64 `json:"ware_house_time,omitempty" xml:"ware_house_time,omitempty"`
    

    // 商品业务标签，不可修改
    
    ItemTags   []string `json:"item_tags,omitempty" xml:"item_tags>string,omitempty"`
    

    // 渠道类目id，不可修改
    
    ChannelCatId   string `json:"channel_cat_id,omitempty" xml:"channel_cat_id,omitempty"`
    

    // 业务模式 0 已验货不入仓，1 已验货入仓，2 普通商品
    
    ItemBizType   string `json:"item_biz_type,omitempty" xml:"item_biz_type,omitempty"`
    

}
