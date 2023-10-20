package smartstore

import (
	"sync"
)

// TmallPopupstoreActivityDeviceQueryResultDto 结构体
type TmallPopupstoreActivityDeviceQueryResultDto struct {
	// 结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 返回结果
	Result *TmallPopupstoreActivityDeviceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTmallPopupstoreActivityDeviceQueryResultDto = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityDeviceQueryResultDto)
	},
}

// GetTmallPopupstoreActivityDeviceQueryResultDto() 从对象池中获取TmallPopupstoreActivityDeviceQueryResultDto
func GetTmallPopupstoreActivityDeviceQueryResultDto() *TmallPopupstoreActivityDeviceQueryResultDto {
	return poolTmallPopupstoreActivityDeviceQueryResultDto.Get().(*TmallPopupstoreActivityDeviceQueryResultDto)
}

// ReleaseTmallPopupstoreActivityDeviceQueryResultDto 释放TmallPopupstoreActivityDeviceQueryResultDto
func ReleaseTmallPopupstoreActivityDeviceQueryResultDto(v *TmallPopupstoreActivityDeviceQueryResultDto) {
	v.Code = ""
	v.Msg = ""
	v.Total = 0
	v.Result = nil
	poolTmallPopupstoreActivityDeviceQueryResultDto.Put(v)
}
