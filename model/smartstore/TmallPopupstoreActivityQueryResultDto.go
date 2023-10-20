package smartstore

import (
	"sync"
)

// TmallPopupstoreActivityQueryResultDto 结构体
type TmallPopupstoreActivityQueryResultDto struct {
	// 返回结果
	ResultList []TmallPopupstoreActivityQueryResult `json:"result_list,omitempty" xml:"result_list>tmall_popupstore_activity_query_result,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolTmallPopupstoreActivityQueryResultDto = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityQueryResultDto)
	},
}

// GetTmallPopupstoreActivityQueryResultDto() 从对象池中获取TmallPopupstoreActivityQueryResultDto
func GetTmallPopupstoreActivityQueryResultDto() *TmallPopupstoreActivityQueryResultDto {
	return poolTmallPopupstoreActivityQueryResultDto.Get().(*TmallPopupstoreActivityQueryResultDto)
}

// ReleaseTmallPopupstoreActivityQueryResultDto 释放TmallPopupstoreActivityQueryResultDto
func ReleaseTmallPopupstoreActivityQueryResultDto(v *TmallPopupstoreActivityQueryResultDto) {
	v.ResultList = v.ResultList[:0]
	v.Code = ""
	v.Msg = ""
	v.Total = 0
	poolTmallPopupstoreActivityQueryResultDto.Put(v)
}
