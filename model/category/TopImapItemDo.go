package category

// TopImapItemDo 
/* model for simplify = false
type TopImapItemDo struct {

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 【必填】目标渠道ID
    
    TargetChannelId   int64 `json:"target_channel_id,omitempty"`
    

    // 【非必填，能填则填】商品品牌名称
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // 【非必填，能填则填】商品所在叶子类目类目路径名称列表，从一级到叶子
    
    SrcCatNamePathList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"src_cat_name_path_list,omitempty"`
    

    // 【必填】源渠道ID
    
    SrcChannelId   int64 `json:"src_channel_id,omitempty"`
    

    // 【非必填，能填则填】商品相关pv信息
    
    PvPairDoList  struct {
        TopPVPairDO  []TopPVPairDO `json:"top_pv_pair_do,omitempty"`
    } `json:"pv_pair_do_list,omitempty"`
    

    // 【必填】商品标题
    
    Title   string `json:"title,omitempty"`
    

    // 【非必填】
    
    TargetCategoryId   int64 `json:"target_category_id,omitempty"`
    

    // 【必填】商品所在叶子类目ID
    
    SrcCategoryId   int64 `json:"src_category_id,omitempty"`
    

    // 非必填，能填则填】barccode列表
    
    BarcodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"barcode_list,omitempty"`
    

}
*/

// TopImapItemDo 
type TopImapItemDo struct {

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 【必填】目标渠道ID
    TargetChannelId   int64 `json:"target_channel_id,omitempty"`

    // 【非必填，能填则填】商品品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // 【非必填，能填则填】商品所在叶子类目类目路径名称列表，从一级到叶子
    SrcCatNamePathList   []string `json:"src_cat_name_path_list,omitempty"`

    // 【必填】源渠道ID
    SrcChannelId   int64 `json:"src_channel_id,omitempty"`

    // 【非必填，能填则填】商品相关pv信息
    PvPairDoList   []TopPVPairDO `json:"pv_pair_do_list,omitempty"`

    // 【必填】商品标题
    Title   string `json:"title,omitempty"`

    // 【非必填】
    TargetCategoryId   int64 `json:"target_category_id,omitempty"`

    // 【必填】商品所在叶子类目ID
    SrcCategoryId   int64 `json:"src_category_id,omitempty"`

    // 非必填，能填则填】barccode列表
    BarcodeList   []string `json:"barcode_list,omitempty"`

}
