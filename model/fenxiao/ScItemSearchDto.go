package fenxiao

// ScItemSearchDto 结构体
type ScItemSearchDto struct {
	// 剔除的货品类型
	NonScItemTypeList []string `json:"non_sc_item_type_list,omitempty" xml:"non_sc_item_type_list>string,omitempty"`
	// 货品类型
	ScItemTypeList []string `json:"sc_item_type_list,omitempty" xml:"sc_item_type_list>string,omitempty"`
	// 货品编码列表
	OuterIdList []string `json:"outer_id_list,omitempty" xml:"outer_id_list>string,omitempty"`
	// 条形码列表
	WhcBarCodeList []string `json:"whc_bar_code_list,omitempty" xml:"whc_bar_code_list>string,omitempty"`
	// 商家类目ID列表
	MerchantCategoryIdList []int64 `json:"merchant_category_id_list,omitempty" xml:"merchant_category_id_list>int64,omitempty"`
	// 标签
	AuctionTags []string `json:"auction_tags,omitempty" xml:"auction_tags>string,omitempty"`
	// 货品ID列表
	ScItemIdList []int64 `json:"sc_item_id_list,omitempty" xml:"sc_item_id_list>int64,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 货品不包含的特征
	NonFeatureCondition string `json:"non_feature_condition,omitempty" xml:"non_feature_condition,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品包含的特征
	FeatureCondition string `json:"feature_condition,omitempty" xml:"feature_condition,omitempty"`
	// 条形码
	WhcBarcode string `json:"whc_barcode,omitempty" xml:"whc_barcode,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 标记位，3表示后端商品
	ItemDim int64 `json:"item_dim,omitempty" xml:"item_dim,omitempty"`
	// 货品状态，0 正常
	ScItemStatus int64 `json:"sc_item_status,omitempty" xml:"sc_item_status,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 市场类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
