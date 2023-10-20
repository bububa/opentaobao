package product

import (
	"sync"
)

// TmallItemStoreUpdateSchemaGetApiResult 结构体
type TmallItemStoreUpdateSchemaGetApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 商品标题，价格，图片等商品数据的schema xml
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误信息
	MappedErrorMessages string `json:"mapped_error_messages,omitempty" xml:"mapped_error_messages,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTmallItemStoreUpdateSchemaGetApiResult = sync.Pool{
	New: func() any {
		return new(TmallItemStoreUpdateSchemaGetApiResult)
	},
}

// GetTmallItemStoreUpdateSchemaGetApiResult() 从对象池中获取TmallItemStoreUpdateSchemaGetApiResult
func GetTmallItemStoreUpdateSchemaGetApiResult() *TmallItemStoreUpdateSchemaGetApiResult {
	return poolTmallItemStoreUpdateSchemaGetApiResult.Get().(*TmallItemStoreUpdateSchemaGetApiResult)
}

// ReleaseTmallItemStoreUpdateSchemaGetApiResult 释放TmallItemStoreUpdateSchemaGetApiResult
func ReleaseTmallItemStoreUpdateSchemaGetApiResult(v *TmallItemStoreUpdateSchemaGetApiResult) {
	v.ErMsg = ""
	v.ErCode = ""
	v.Result = ""
	v.MappedErrorMessages = ""
	v.Error = false
	poolTmallItemStoreUpdateSchemaGetApiResult.Put(v)
}
