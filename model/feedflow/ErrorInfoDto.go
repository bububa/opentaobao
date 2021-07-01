package feedflow

// ErrorInfoDto 结构体
type ErrorInfoDto struct {
	// 失败原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 该原因失败对象列表
	ErrorObjectList []ErrorObjectDto `json:"error_object_list,omitempty" xml:"error_object_list>error_object_dto,omitempty"`
}
