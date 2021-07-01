package wdk

// ContainerDto 结构体
type ContainerDto struct {
	// 容器code
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 容器类型
	ContainerType string `json:"container_type,omitempty" xml:"container_type,omitempty"`
}
