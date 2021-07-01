package category

// TopImapItemDo 
type TopImapItemDo struct {
    // 商品ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 【必填】目标渠道ID
    TargetChannelId   int64 `json:"target_channel_id,omitempty" xml:"target_channel_id,omitempty"`
    // 【非必填，能填则填】商品品牌名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 【非必填，能填则填】商品所在叶子类目类目路径名称列表，从一级到叶子
    SrcCatNamePathList   []string `json:"src_cat_name_path_list,omitempty" xml:"src_cat_name_path_list>string,omitempty"`
    // 【必填】源渠道ID
    SrcChannelId   int64 `json:"src_channel_id,omitempty" xml:"src_channel_id,omitempty"`
    // 【非必填，能填则填】商品相关pv信息
    PvPairDoList   []TopPVPairDo `json:"pv_pair_do_list,omitempty" xml:"pv_pair_do_list>top_pv_pair_do,omitempty"`
    // 【必填】商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 【非必填】
    TargetCategoryId   int64 `json:"target_category_id,omitempty" xml:"target_category_id,omitempty"`
    // 【必填】商品所在叶子类目ID
    SrcCategoryId   int64 `json:"src_category_id,omitempty" xml:"src_category_id,omitempty"`
    // 非必填，能填则填】barccode列表
    BarcodeList   []string `json:"barcode_list,omitempty" xml:"barcode_list>string,omitempty"`
}
