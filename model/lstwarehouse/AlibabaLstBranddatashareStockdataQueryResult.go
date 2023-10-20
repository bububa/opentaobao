package lstwarehouse

import (
	"sync"
)

// AlibabaLstBranddatashareStockdataQueryResult 结构体
type AlibabaLstBranddatashareStockdataQueryResult struct {
	// 返回内容
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstBranddatashareStockdataQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstBranddatashareStockdataQueryResult)
	},
}

// GetAlibabaLstBranddatashareStockdataQueryResult() 从对象池中获取AlibabaLstBranddatashareStockdataQueryResult
func GetAlibabaLstBranddatashareStockdataQueryResult() *AlibabaLstBranddatashareStockdataQueryResult {
	return poolAlibabaLstBranddatashareStockdataQueryResult.Get().(*AlibabaLstBranddatashareStockdataQueryResult)
}

// ReleaseAlibabaLstBranddatashareStockdataQueryResult 释放AlibabaLstBranddatashareStockdataQueryResult
func ReleaseAlibabaLstBranddatashareStockdataQueryResult(v *AlibabaLstBranddatashareStockdataQueryResult) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Total = 0
	v.Size = 0
	v.Page = 0
	v.Success = false
	poolAlibabaLstBranddatashareStockdataQueryResult.Put(v)
}
