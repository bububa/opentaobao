package cainiaolocker

import (
	"sync"
)

// WaybillCloudPrintNewResponse 结构体
type WaybillCloudPrintNewResponse struct {
	// 云打印内容（encryptedData表示加密结果，data表示非加密结果）;模板内容,具体解释见&lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.8cf9Nj&amp;treeId=17&amp;articleId=105085&amp;docType=1#12&#34;&gt;链接&lt;/a&gt;
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
	// 面单号, 子母件模式下为子面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 子母件中的母单号，当为子母件模式时，需要此单号为实际挂载物流详情的单号，需要使用此单号进行发货，查询物流详情，非子母件，此字段为空
	ParentWaybillCode string `json:"parent_waybill_code,omitempty" xml:"parent_waybill_code,omitempty"`
	// 拓展信息，特殊场景下使用
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// cp_code，跟入参传入的cp_code保持一致
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 真实取号的cp_code，在菜鸟裹裹商家寄件等虚拟cp的场景中real_cp_code和入参中的cp_code不一样，其他场景二者一样
	RealCpCode string `json:"real_cp_code,omitempty" xml:"real_cp_code,omitempty"`
	// 请求id
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 本单请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolWaybillCloudPrintNewResponse = sync.Pool{
	New: func() any {
		return new(WaybillCloudPrintNewResponse)
	},
}

// GetWaybillCloudPrintNewResponse() 从对象池中获取WaybillCloudPrintNewResponse
func GetWaybillCloudPrintNewResponse() *WaybillCloudPrintNewResponse {
	return poolWaybillCloudPrintNewResponse.Get().(*WaybillCloudPrintNewResponse)
}

// ReleaseWaybillCloudPrintNewResponse 释放WaybillCloudPrintNewResponse
func ReleaseWaybillCloudPrintNewResponse(v *WaybillCloudPrintNewResponse) {
	v.PrintData = ""
	v.WaybillCode = ""
	v.ParentWaybillCode = ""
	v.ExtraInfo = ""
	v.CpCode = ""
	v.RealCpCode = ""
	v.ObjectId = ""
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.IsSuccess = false
	poolWaybillCloudPrintNewResponse.Put(v)
}
