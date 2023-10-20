package category

import (
	"sync"
)

// AlibabaImapFixedmappingQueryResult 结构体
type AlibabaImapFixedmappingQueryResult struct {
	// list参数
	TopImapUnionCategoryPathDoList []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do_list,omitempty" xml:"top_imap_union_category_path_do_list>top_imap_union_category_path_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaImapFixedmappingQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaImapFixedmappingQueryResult)
	},
}

// GetAlibabaImapFixedmappingQueryResult() 从对象池中获取AlibabaImapFixedmappingQueryResult
func GetAlibabaImapFixedmappingQueryResult() *AlibabaImapFixedmappingQueryResult {
	return poolAlibabaImapFixedmappingQueryResult.Get().(*AlibabaImapFixedmappingQueryResult)
}

// ReleaseAlibabaImapFixedmappingQueryResult 释放AlibabaImapFixedmappingQueryResult
func ReleaseAlibabaImapFixedmappingQueryResult(v *AlibabaImapFixedmappingQueryResult) {
	v.TopImapUnionCategoryPathDoList = v.TopImapUnionCategoryPathDoList[:0]
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaImapFixedmappingQueryResult.Put(v)
}
