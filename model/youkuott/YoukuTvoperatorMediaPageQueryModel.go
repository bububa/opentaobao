package youkuott

import (
	"sync"
)

// YoukuTvoperatorMediaPageQueryModel 结构体
type YoukuTvoperatorMediaPageQueryModel struct {
	// 数据列表
	DataList []YoukuTvoperatorMediaPageQueryData `json:"data_list,omitempty" xml:"data_list>youku_tvoperator_media_page_query_data,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 页号
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolYoukuTvoperatorMediaPageQueryModel = sync.Pool{
	New: func() any {
		return new(YoukuTvoperatorMediaPageQueryModel)
	},
}

// GetYoukuTvoperatorMediaPageQueryModel() 从对象池中获取YoukuTvoperatorMediaPageQueryModel
func GetYoukuTvoperatorMediaPageQueryModel() *YoukuTvoperatorMediaPageQueryModel {
	return poolYoukuTvoperatorMediaPageQueryModel.Get().(*YoukuTvoperatorMediaPageQueryModel)
}

// ReleaseYoukuTvoperatorMediaPageQueryModel 释放YoukuTvoperatorMediaPageQueryModel
func ReleaseYoukuTvoperatorMediaPageQueryModel(v *YoukuTvoperatorMediaPageQueryModel) {
	v.DataList = v.DataList[:0]
	v.TotalCount = 0
	v.PageNo = 0
	v.PageSize = 0
	v.TotalPage = 0
	v.HasNext = false
	poolYoukuTvoperatorMediaPageQueryModel.Put(v)
}
