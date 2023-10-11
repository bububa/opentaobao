package simba

// StdCategoryQueryVo 结构体
type StdCategoryQueryVo struct {
	// 宝贝id集合
	MaterialIdList []int64 `json:"material_id_list,omitempty" xml:"material_id_list>int64,omitempty"`
	// 推广主体类型,item:商品,item_private_mini:独享橱窗,shop:店铺,content:内容,short_video:短视频,user_define:自定义;
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 父级类目级别
	ParentCatLevel int64 `json:"parent_cat_level,omitempty" xml:"parent_cat_level,omitempty"`
}
