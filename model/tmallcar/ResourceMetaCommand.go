package tmallcar

import (
	"sync"
)

// ResourceMetaCommand 结构体
type ResourceMetaCommand struct {
	// 摘要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 扩展字段
	ExtensionField string `json:"extension_field,omitempty" xml:"extension_field,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 来源
	FromSource string `json:"from_source,omitempty" xml:"from_source,omitempty"`
	// 资源类型
	ResourceType string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 资源附属信息
	ResourceOptions string `json:"resource_options,omitempty" xml:"resource_options,omitempty"`
	// 资源id
	ResourceId int64 `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 外部车系id
	OutSeriesId int64 `json:"out_series_id,omitempty" xml:"out_series_id,omitempty"`
	// 分组id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 外部车型id
	OutModelId int64 `json:"out_model_id,omitempty" xml:"out_model_id,omitempty"`
}

var poolResourceMetaCommand = sync.Pool{
	New: func() any {
		return new(ResourceMetaCommand)
	},
}

// GetResourceMetaCommand() 从对象池中获取ResourceMetaCommand
func GetResourceMetaCommand() *ResourceMetaCommand {
	return poolResourceMetaCommand.Get().(*ResourceMetaCommand)
}

// ReleaseResourceMetaCommand 释放ResourceMetaCommand
func ReleaseResourceMetaCommand(v *ResourceMetaCommand) {
	v.Summary = ""
	v.ExtensionField = ""
	v.Title = ""
	v.FromSource = ""
	v.ResourceType = ""
	v.ResourceOptions = ""
	v.ResourceId = 0
	v.OutSeriesId = 0
	v.GroupId = 0
	v.OutModelId = 0
	poolResourceMetaCommand.Put(v)
}
