package ju

// Items 
/* model for simplify = false
type Items struct {

    // 卖点描述
    
    UspDescList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"usp_desc_list,omitempty"`
    

    // 淘宝类目id
    
    TbFirstCatId   int64 `json:"tb_first_cat_id,omitempty"`
    

    // 原价
    
    OrigPrice   string `json:"orig_price,omitempty"`
    

    // itemId
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 展示结束时间
    
    ShowEndTime   int64 `json:"show_end_time,omitempty"`
    

    // pc链接
    
    PcUrl   string `json:"pc_url,omitempty"`
    

    // 频道id
    
    PlatformId   int64 `json:"platform_id,omitempty"`
    

    // 聚划算id
    
    JuId   int64 `json:"ju_id,omitempty"`
    

    // 无线主图
    
    PicUrlForWL   string `json:"pic_url_for_w_l,omitempty"`
    

    // 开团时间
    
    OnlineStartTime   int64 `json:"online_start_time,omitempty"`
    

    // 类目名称
    
    CategoryName   string `json:"category_name,omitempty"`
    

    // 聚划算价格，单位分
    
    ActPrice   string `json:"act_price,omitempty"`
    

    // 商品标题
    
    Title   string `json:"title,omitempty"`
    

    // 无线链接
    
    WapUrl   string `json:"wap_url,omitempty"`
    

    // 商品卖点
    
    ItemUspList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"item_usp_list,omitempty"`
    

    // 开始展示时间
    
    ShowStartTime   int64 `json:"show_start_time,omitempty"`
    

    // 开团结束时间
    
    OnlineEndTime   int64 `json:"online_end_time,omitempty"`
    

    // pc主图
    
    PicUrlForPC   string `json:"pic_url_for_p_c,omitempty"`
    

    // 价格卖点
    
    PriceUspList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"price_usp_list,omitempty"`
    

    // 是否包邮
    
    PayPostage   bool `json:"pay_postage,omitempty"`
    

}
*/

// Items 
type Items struct {

    // 卖点描述
    UspDescList   []string `json:"usp_desc_list,omitempty"`

    // 淘宝类目id
    TbFirstCatId   int64 `json:"tb_first_cat_id,omitempty"`

    // 原价
    OrigPrice   string `json:"orig_price,omitempty"`

    // itemId
    ItemId   int64 `json:"item_id,omitempty"`

    // 展示结束时间
    ShowEndTime   int64 `json:"show_end_time,omitempty"`

    // pc链接
    PcUrl   string `json:"pc_url,omitempty"`

    // 频道id
    PlatformId   int64 `json:"platform_id,omitempty"`

    // 聚划算id
    JuId   int64 `json:"ju_id,omitempty"`

    // 无线主图
    PicUrlForWL   string `json:"pic_url_for_w_l,omitempty"`

    // 开团时间
    OnlineStartTime   int64 `json:"online_start_time,omitempty"`

    // 类目名称
    CategoryName   string `json:"category_name,omitempty"`

    // 聚划算价格，单位分
    ActPrice   string `json:"act_price,omitempty"`

    // 商品标题
    Title   string `json:"title,omitempty"`

    // 无线链接
    WapUrl   string `json:"wap_url,omitempty"`

    // 商品卖点
    ItemUspList   []string `json:"item_usp_list,omitempty"`

    // 开始展示时间
    ShowStartTime   int64 `json:"show_start_time,omitempty"`

    // 开团结束时间
    OnlineEndTime   int64 `json:"online_end_time,omitempty"`

    // pc主图
    PicUrlForPC   string `json:"pic_url_for_p_c,omitempty"`

    // 价格卖点
    PriceUspList   []string `json:"price_usp_list,omitempty"`

    // 是否包邮
    PayPostage   bool `json:"pay_postage,omitempty"`

}
