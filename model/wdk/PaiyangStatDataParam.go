package wdk

// PaiyangStatDataParam 
type PaiyangStatDataParam struct {

    // 经营店编码
    ShopCode   string `json:"shop_code,omitempty"`

    // 统计时间
    StatDate   string `json:"stat_date,omitempty"`

    // 分页页码
    Current   int64 `json:"current,omitempty"`

    // 活动id集合，最大支持20个
    ActivityIdList   []String `json:"activity_id_list,omitempty"`

    // 69码集合，最大支持20个
    BarcodeList   []String `json:"barcode_list,omitempty"`

    // 分页页大小
    PageSize   int64 `json:"page_size,omitempty"`

}
