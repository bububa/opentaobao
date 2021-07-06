package campus

// TreeNode 结构体
type TreeNode struct {
	// childs
	Childs []string `json:"childs,omitempty" xml:"childs>string,omitempty"`
	// 菜单主键id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 菜单名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 菜单链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 菜单层级路劲
	MenuId string `json:"menu_id,omitempty" xml:"menu_id,omitempty"`
	// 菜单级别
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 菜单排序字段
	MenuOrder int64 `json:"menu_order,omitempty" xml:"menu_order,omitempty"`
}
