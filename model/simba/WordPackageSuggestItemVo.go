package simba

// WordPackageSuggestItemVo 结构体
type WordPackageSuggestItemVo struct {
	// 推荐的词包列表
	WordPackageList []SuggestWordPackageVo `json:"word_package_list,omitempty" xml:"word_package_list>suggest_word_package_vo,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}
