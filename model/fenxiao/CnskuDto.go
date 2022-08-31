package fenxiao

// CnskuDto 结构体
type CnskuDto struct {
	// 组合货品
	CombRelationDTOList []CnskuRelationDto `json:"comb_relation_d_t_o_list,omitempty" xml:"comb_relation_d_t_o_list>cnsku_relation_dto,omitempty"`
	// 套装货品关系列表
	SuiteGoodsRelationList []CnskuSuiteGoodsRelationDto `json:"suite_goods_relation_list,omitempty" xml:"suite_goods_relation_list>cnsku_suite_goods_relation_dto,omitempty"`
	// 需要向featureMap中增加的属性列表（新增&amp;更新接口使用）
	UpdateFeatureMap string `json:"update_feature_map,omitempty" xml:"update_feature_map,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品类型，NORMAL / COMBINE / SUITEGOODS
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 条形码( 格式如：条码1#条码2#条码3，多条码中间用 # 分隔)
	WhcBarCode string `json:"whc_bar_code,omitempty" xml:"whc_bar_code,omitempty"`
	// 货品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 想去除的feature 和 updateFeatureMap对应
	RemoveFeatureMap string `json:"remove_feature_map,omitempty" xml:"remove_feature_map,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 包装材料
	PackageMaterial string `json:"package_material,omitempty" xml:"package_material,omitempty"`
	// 扩展属性
	FeatureMap string `json:"feature_map,omitempty" xml:"feature_map,omitempty"`
	// 货品高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 货品重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 吊牌价
	ReservePrice int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 扩展字段
	CnskuExtendDTO *CnskuExtendDto `json:"cnsku_extend_d_t_o,omitempty" xml:"cnsku_extend_d_t_o,omitempty"`
	// feature字段
	CnskuFeatureDTO *CnskuFeatureDto `json:"cnsku_feature_d_t_o,omitempty" xml:"cnsku_feature_d_t_o,omitempty"`
	// 长度
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 市场叶子类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 货品Id
	CnskuId int64 `json:"cnsku_id,omitempty" xml:"cnsku_id,omitempty"`
	// 货主id
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	// 前端skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 货品分类
	SkuCategory int64 `json:"sku_category,omitempty" xml:"sku_category,omitempty"`
	// 体积
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
