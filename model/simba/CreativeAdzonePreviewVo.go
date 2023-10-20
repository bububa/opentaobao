package simba

import (
	"sync"
)

// CreativeAdzonePreviewVo 结构体
type CreativeAdzonePreviewVo struct {
	// 创意预览返回前端展示对象
	PreviewDataList []CreativePreviewResultVo `json:"preview_data_list,omitempty" xml:"preview_data_list>creative_preview_result_vo,omitempty"`
	// 资源包名字
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 资源位类型,1:首焦,2:猜你喜欢,3:NewDetail
	AdzoneType int64 `json:"adzone_type,omitempty" xml:"adzone_type,omitempty"`
	// 资源包id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
}

var poolCreativeAdzonePreviewVo = sync.Pool{
	New: func() any {
		return new(CreativeAdzonePreviewVo)
	},
}

// GetCreativeAdzonePreviewVo() 从对象池中获取CreativeAdzonePreviewVo
func GetCreativeAdzonePreviewVo() *CreativeAdzonePreviewVo {
	return poolCreativeAdzonePreviewVo.Get().(*CreativeAdzonePreviewVo)
}

// ReleaseCreativeAdzonePreviewVo 释放CreativeAdzonePreviewVo
func ReleaseCreativeAdzonePreviewVo(v *CreativeAdzonePreviewVo) {
	v.PreviewDataList = v.PreviewDataList[:0]
	v.AdzoneName = ""
	v.AdzoneType = 0
	v.AdzoneId = 0
	poolCreativeAdzonePreviewVo.Put(v)
}
