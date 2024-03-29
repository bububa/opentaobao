package media

import (
	"sync"
)

// SearchResultDo 结构体
type SearchResultDo struct {
	// 视频信息列表
	ResultList []VideoItemExtDo `json:"result_list,omitempty" xml:"result_list>video_item_ext_do,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 总视频数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolSearchResultDo = sync.Pool{
	New: func() any {
		return new(SearchResultDo)
	},
}

// GetSearchResultDo() 从对象池中获取SearchResultDo
func GetSearchResultDo() *SearchResultDo {
	return poolSearchResultDo.Get().(*SearchResultDo)
}

// ReleaseSearchResultDo 释放SearchResultDo
func ReleaseSearchResultDo(v *SearchResultDo) {
	v.ResultList = v.ResultList[:0]
	v.ResultCode = ""
	v.TotalNum = 0
	v.IsSuccess = false
	poolSearchResultDo.Put(v)
}
