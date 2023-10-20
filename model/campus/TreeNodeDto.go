package campus

import (
	"sync"
)

// TreeNodeDto 结构体
type TreeNodeDto struct {
	// data
	Datas []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
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
}

var poolTreeNodeDto = sync.Pool{
	New: func() any {
		return new(TreeNodeDto)
	},
}

// GetTreeNodeDto() 从对象池中获取TreeNodeDto
func GetTreeNodeDto() *TreeNodeDto {
	return poolTreeNodeDto.Get().(*TreeNodeDto)
}

// ReleaseTreeNodeDto 释放TreeNodeDto
func ReleaseTreeNodeDto(v *TreeNodeDto) {
	v.Datas = v.Datas[:0]
	v.Id = ""
	v.Pid = ""
	v.Url = ""
	v.Name = ""
	v.NodeType = ""
	v.DataType = ""
	poolTreeNodeDto.Put(v)
}
