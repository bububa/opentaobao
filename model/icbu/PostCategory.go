package icbu

// PostCategory 
/* model for simplify = false
type PostCategory struct {

    // 父类目ID数组
    
    ParentIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"parent_ids,omitempty"`
    

    // 类目层级
    
    Level   int64 `json:"level,omitempty"`
    

    // 是否叶子类目（只有叶子类目才能发布商品）
    
    LeafCategory   bool `json:"leaf_category,omitempty"`
    

    // 类目名称
    
    Name   string `json:"name,omitempty"`
    

    // 类目ID
    
    CategoryId   int64 `json:"category_id,omitempty"`
    

    // 子类目ID数组
    
    ChildIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"child_ids,omitempty"`
    

    // 类目的中文名称
    
    CnName   string `json:"cn_name,omitempty"`
    

}
*/

// PostCategory 
type PostCategory struct {

    // 父类目ID数组
    ParentIds   []int64 `json:"parent_ids,omitempty"`

    // 类目层级
    Level   int64 `json:"level,omitempty"`

    // 是否叶子类目（只有叶子类目才能发布商品）
    LeafCategory   bool `json:"leaf_category,omitempty"`

    // 类目名称
    Name   string `json:"name,omitempty"`

    // 类目ID
    CategoryId   int64 `json:"category_id,omitempty"`

    // 子类目ID数组
    ChildIds   []int64 `json:"child_ids,omitempty"`

    // 类目的中文名称
    CnName   string `json:"cn_name,omitempty"`

}
