package campus

// TreeNodeDto 结构体
type TreeNodeDto struct {
	// 权限id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 父id
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 权限名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点类型
	NodeType string `json:"node_type,omitempty" xml:"node_type,omitempty"`
	// 权限类型
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// data
	Datas []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
}
