package icbu

// ProductGroup 
type ProductGroup struct {

    // 上级分组ID
    ParentId   int64 `json:"parent_id,omitempty"`

    // 分组ID
    GroupId   int64 `json:"group_id,omitempty"`

    // 分组名称
    GroupName   string `json:"group_name,omitempty"`

    // 下级分组ID列表
    ChildrenIdList   []Number `json:"children_id_list,omitempty"`

    // 父节点id，父节点处在分组树的二级
    ParentId2   int64 `json:"parent_id2,omitempty"`

}