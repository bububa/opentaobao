package wdk

// PaiyangStatDataParam 
type PaiyangStatDataParam struct {

    // 经营店编码
    
    ShopCode   string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
    

    // 统计时间
    
    StatDate   string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
    

    // 分页页码
    
    Current   int64 `json:"current,omitempty" xml:"current,omitempty"`
    

    // 活动id集合，最大支持20个
    
    ActivityIdList   []string `json:"activity_id_list,omitempty" xml:"activity_id_list>string,omitempty"`
    

    // 69码集合，最大支持20个
    
    BarcodeList   []string `json:"barcode_list,omitempty" xml:"barcode_list>string,omitempty"`
    

    // 分页页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
