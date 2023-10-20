package servicecenter

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarContractDownloadResult 结构体
type TmallCarContractDownloadResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 当前时间
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// 返回的合同二进制文件
	Objects *model.File `json:"objects,omitempty" xml:"objects,omitempty"`
	// 消耗时间
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// 成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallCarContractDownloadResult = sync.Pool{
	New: func() any {
		return new(TmallCarContractDownloadResult)
	},
}

// GetTmallCarContractDownloadResult() 从对象池中获取TmallCarContractDownloadResult
func GetTmallCarContractDownloadResult() *TmallCarContractDownloadResult {
	return poolTmallCarContractDownloadResult.Get().(*TmallCarContractDownloadResult)
}

// ReleaseTmallCarContractDownloadResult 释放TmallCarContractDownloadResult
func ReleaseTmallCarContractDownloadResult(v *TmallCarContractDownloadResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.GmtCurrentTime = 0
	v.Objects = nil
	v.CostTime = 0
	v.Success = false
	poolTmallCarContractDownloadResult.Put(v)
}
