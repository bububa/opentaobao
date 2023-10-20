package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSyncedorderQueryAPIResponse 五道口查询同步订单 API返回值
// alibaba.wdk.syncedorder.query
//
// 外部商户查询同步到五道口的订单
type AlibabaWdkSyncedorderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSyncedorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSyncedorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSyncedorderQueryAPIResponseModel).Reset()
}

// AlibabaWdkSyncedorderQueryAPIResponseModel is 五道口查询同步订单 成功返回结果
type AlibabaWdkSyncedorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_syncedorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 返回码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSyncedorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.BizOrderId = ""
	m.ReturnCode = 0
	m.IsSuccess = false
}

var poolAlibabaWdkSyncedorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSyncedorderQueryAPIResponse)
	},
}

// GetAlibabaWdkSyncedorderQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSyncedorderQueryAPIResponse
func GetAlibabaWdkSyncedorderQueryAPIResponse() *AlibabaWdkSyncedorderQueryAPIResponse {
	return poolAlibabaWdkSyncedorderQueryAPIResponse.Get().(*AlibabaWdkSyncedorderQueryAPIResponse)
}

// ReleaseAlibabaWdkSyncedorderQueryAPIResponse 将 AlibabaWdkSyncedorderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSyncedorderQueryAPIResponse(v *AlibabaWdkSyncedorderQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSyncedorderQueryAPIResponse.Put(v)
}
