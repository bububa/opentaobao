package tmallcampus

import (
	"sync"
)

// TmallCampusAuthstatusQueryResult 结构体
type TmallCampusAuthstatusQueryResult struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果总数（用于分页）
	TotalNum string `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 页码（用于分页）
	PageNum string `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 业务数据-学生认证状态
	Data *StudentDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTmallCampusAuthstatusQueryResult = sync.Pool{
	New: func() any {
		return new(TmallCampusAuthstatusQueryResult)
	},
}

// GetTmallCampusAuthstatusQueryResult() 从对象池中获取TmallCampusAuthstatusQueryResult
func GetTmallCampusAuthstatusQueryResult() *TmallCampusAuthstatusQueryResult {
	return poolTmallCampusAuthstatusQueryResult.Get().(*TmallCampusAuthstatusQueryResult)
}

// ReleaseTmallCampusAuthstatusQueryResult 释放TmallCampusAuthstatusQueryResult
func ReleaseTmallCampusAuthstatusQueryResult(v *TmallCampusAuthstatusQueryResult) {
	v.Msg = ""
	v.Code = ""
	v.TotalNum = ""
	v.Success = ""
	v.PageNum = ""
	v.Data = nil
	poolTmallCampusAuthstatusQueryResult.Put(v)
}
