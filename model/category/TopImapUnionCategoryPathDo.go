package category

// TopImapUnionCategoryPathDo 结构体
type TopImapUnionCategoryPathDo struct {
	// 叶子类目名称
	LeafCatName string `json:"leaf_cat_name,omitempty" xml:"leaf_cat_name,omitempty"`
	// 四级类目名称
	L4CatName string `json:"l4_cat_name,omitempty" xml:"l4_cat_name,omitempty"`
	// 二级类目名称
	L2CatName string `json:"l2_cat_name,omitempty" xml:"l2_cat_name,omitempty"`
	// 三级类目名称
	L3CatName string `json:"l3_cat_name,omitempty" xml:"l3_cat_name,omitempty"`
	// 五级类目名称
	L5CatName string `json:"l5_cat_name,omitempty" xml:"l5_cat_name,omitempty"`
	// 一级类目名称
	L1CatName string `json:"l1_cat_name,omitempty" xml:"l1_cat_name,omitempty"`
	// 五级类目ID
	L5CatId int64 `json:"l5_cat_id,omitempty" xml:"l5_cat_id,omitempty"`
	// 三级类目ID
	L3CatId int64 `json:"l3_cat_id,omitempty" xml:"l3_cat_id,omitempty"`
	// 叶子类目ID
	LeafCatId int64 `json:"leaf_cat_id,omitempty" xml:"leaf_cat_id,omitempty"`
	// 一级类目ID
	L1CatId int64 `json:"l1_cat_id,omitempty" xml:"l1_cat_id,omitempty"`
	// 四级类目ID
	L4CatId int64 `json:"l4_cat_id,omitempty" xml:"l4_cat_id,omitempty"`
	// 二级类目ID
	L2CatId int64 `json:"l2_cat_id,omitempty" xml:"l2_cat_id,omitempty"`
	// 渠道ID
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 是否叶子
	Leaf bool `json:"leaf,omitempty" xml:"leaf,omitempty"`
}
