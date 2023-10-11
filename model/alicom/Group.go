package alicom

// Group 结构体
type Group struct {
	// 接口
	InterfaceList []Interface `json:"interface_list,omitempty" xml:"interface_list>interface,omitempty"`
	// spu的结构
	SpuMap string `json:"spu_map,omitempty" xml:"spu_map,omitempty"`
	// 是否主接口
	Main bool `json:"main,omitempty" xml:"main,omitempty"`
}
