package wdk

// ApiScrollPageResult 
type ApiScrollPageResult struct {
    // 条目总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 第一次调用会返回scroll_id 但并不包含数据，后续每次都必须指定上一次返回scroll_id，并且后续搜索结果中都会返回scroll_id及对应匹配的数据，后续查询该参数必填
    ScrollId   string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
    // 结果页数
    PageCount   int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
    // 商品主档对象
    ModelList   []WdkOpenMerchantSkuDO `json:"model_list,omitempty" xml:"model_list>wdk_open_merchant_sku_do,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 当前页码
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    // 数据集合
    Models   []WarehouseSkuDO `json:"models,omitempty" xml:"models>warehouse_sku_do,omitempty"`
}
