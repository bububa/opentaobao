package crm

// Group 
type Group struct {

    // 分组的id
    
    GroupId   int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
    

    // 分组的名称
    
    GroupName   string `json:"group_name,omitempty" xml:"group_name,omitempty"`
    

    // 分组的创建时间
    
    GroupCreate   string `json:"group_create,omitempty" xml:"group_create,omitempty"`
    

    // 分组的修改时间
    
    GroupModify   string `json:"group_modify,omitempty" xml:"group_modify,omitempty"`
    

    // 分组的状态，1表示正常
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 分组所拥有的会员数量,如果返回值为-1，表示当前服务忙或数据在初始化。
    
    MemberCount   int64 `json:"member_count,omitempty" xml:"member_count,omitempty"`
    

}
