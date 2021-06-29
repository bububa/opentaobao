package icbu

// PostCategory 
type PostCategory struct {
    // 父类目ID数组
    ParentIds   []int64 `json:"parent_ids,omitempty" xml:"parent_ids>int64,omitempty"`
    // 类目层级
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    // 是否叶子类目（只有叶子类目才能发布商品）
    LeafCategory   bool `json:"leaf_category,omitempty" xml:"leaf_category,omitempty"`
    // 类目名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 类目ID
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // 子类目ID数组
    ChildIds   []int64 `json:"child_ids,omitempty" xml:"child_ids>int64,omitempty"`
    // 类目的中文名称
    CnName   string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
}
