package alsc

import (
	"sync"
)

// CommonPageResult 结构体
type CommonPageResult struct {
	// 结果
	ResultList []CardTemplateOpenInfo `json:"result_list,omitempty" xml:"result_list>card_template_open_info,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 错误结果显示
	ResultView string `json:"result_view,omitempty" xml:"result_view,omitempty"`
	// 当前页码（原样传出）,若数据下行时为空
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 当前每页显示数量（原样传出）,若数据下行时为空
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数,若数据下行时为空
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
}

var poolCommonPageResult = sync.Pool{
	New: func() any {
		return new(CommonPageResult)
	},
}

// GetCommonPageResult() 从对象池中获取CommonPageResult
func GetCommonPageResult() *CommonPageResult {
	return poolCommonPageResult.Get().(*CommonPageResult)
}

// ReleaseCommonPageResult 释放CommonPageResult
func ReleaseCommonPageResult(v *CommonPageResult) {
	v.ResultList = v.ResultList[:0]
	v.ResultCode = ""
	v.ResultDesc = ""
	v.ResultView = ""
	v.CurrentPage = 0
	v.PageSize = 0
	v.TotalSize = 0
	v.TotalPage = 0
	v.BizSuccess = false
	v.HasNextPage = false
	poolCommonPageResult.Put(v)
}
