package ascpchannel

import (
	"sync"
)

// ResultWrapper 结构体
type ResultWrapper struct {
	// 响应数据
	Datas []Datas `json:"datas,omitempty" xml:"datas>datas,omitempty"`
	// 系统自动生成
	Data []AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData `json:"data,omitempty" xml:"data>alibaba_ascp_aic_supplier_aicinventory_negative_sale_invalidate_data,omitempty"`
	// 返回内容
	DataList []AlibabaAscpChannelMainRefundCreateData `json:"data_list,omitempty" xml:"data_list>alibaba_ascp_channel_main_refund_create_data,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 执行结果
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
	// 是否需要重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolResultWrapper = sync.Pool{
	New: func() any {
		return new(ResultWrapper)
	},
}

// GetResultWrapper() 从对象池中获取ResultWrapper
func GetResultWrapper() *ResultWrapper {
	return poolResultWrapper.Get().(*ResultWrapper)
}

// ReleaseResultWrapper 释放ResultWrapper
func ReleaseResultWrapper(v *ResultWrapper) {
	v.Datas = v.Datas[:0]
	v.Data = v.Data[:0]
	v.DataList = v.DataList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Error = ""
	v.Success = false
	v.ResultSuccess = false
	v.Retry = false
	poolResultWrapper.Put(v)
}
