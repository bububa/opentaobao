package tvupadmin

import (
	"sync"
)

// DfPageResultDto 结构体
type DfPageResultDto struct {
	// value
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// codeName
	CodeName string `json:"code_name,omitempty" xml:"code_name,omitempty"`
	// detailMessage
	DetailMessage string `json:"detail_message,omitempty" xml:"detail_message,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// pageNo
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// totalPage
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDfPageResultDto = sync.Pool{
	New: func() any {
		return new(DfPageResultDto)
	},
}

// GetDfPageResultDto() 从对象池中获取DfPageResultDto
func GetDfPageResultDto() *DfPageResultDto {
	return poolDfPageResultDto.Get().(*DfPageResultDto)
}

// ReleaseDfPageResultDto 释放DfPageResultDto
func ReleaseDfPageResultDto(v *DfPageResultDto) {
	v.Values = v.Values[:0]
	v.CodeName = ""
	v.DetailMessage = ""
	v.Message = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResultCode = ""
	v.Code = 0
	v.PageNo = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.TotalPage = 0
	v.Success = false
	poolDfPageResultDto.Put(v)
}
