package scbp

// TopProductGroupDto 
/* model for simplify = false
type TopProductGroupDto struct {

    // 产品分组标识
    
    GroupId   string `json:"group_id,omitempty"`
    

    // 产品分组名称
    
    GroupName   string `json:"group_name,omitempty"`
    

    // 是否是叶子分组，即没有子分组
    
    Leaf   bool `json:"leaf,omitempty"`
    

}
*/

// TopProductGroupDto 
type TopProductGroupDto struct {

    // 产品分组标识
    GroupId   string `json:"group_id,omitempty"`

    // 产品分组名称
    GroupName   string `json:"group_name,omitempty"`

    // 是否是叶子分组，即没有子分组
    Leaf   bool `json:"leaf,omitempty"`

}
