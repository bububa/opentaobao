package admarket

import (
	"sync"
)

// YunosAdmarketMaterialAuditResult 结构体
type YunosAdmarketMaterialAuditResult struct {
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 返回状态码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolYunosAdmarketMaterialAuditResult = sync.Pool{
	New: func() any {
		return new(YunosAdmarketMaterialAuditResult)
	},
}

// GetYunosAdmarketMaterialAuditResult() 从对象池中获取YunosAdmarketMaterialAuditResult
func GetYunosAdmarketMaterialAuditResult() *YunosAdmarketMaterialAuditResult {
	return poolYunosAdmarketMaterialAuditResult.Get().(*YunosAdmarketMaterialAuditResult)
}

// ReleaseYunosAdmarketMaterialAuditResult 释放YunosAdmarketMaterialAuditResult
func ReleaseYunosAdmarketMaterialAuditResult(v *YunosAdmarketMaterialAuditResult) {
	v.ResultMsg = ""
	v.Result = ""
	v.ResultCode = ""
	v.IsSuccess = false
	poolYunosAdmarketMaterialAuditResult.Put(v)
}
