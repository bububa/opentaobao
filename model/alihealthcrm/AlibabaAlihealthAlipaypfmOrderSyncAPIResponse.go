package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlipaypfmOrderSyncAPIResponse 订单数据回传接口 API返回值
// alibaba.alihealth.alipaypfm.order.sync
//
// 订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计
type AlibabaAlihealthAlipaypfmOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlipaypfmOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel is 订单数据回传接口 成功返回结果
type AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthAlipaypfmOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAlipaypfmOrderSyncAPIResponse)
	},
}

// GetAlibabaAlihealthAlipaypfmOrderSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAlipaypfmOrderSyncAPIResponse
func GetAlibabaAlihealthAlipaypfmOrderSyncAPIResponse() *AlibabaAlihealthAlipaypfmOrderSyncAPIResponse {
	return poolAlibabaAlihealthAlipaypfmOrderSyncAPIResponse.Get().(*AlibabaAlihealthAlipaypfmOrderSyncAPIResponse)
}

// ReleaseAlibabaAlihealthAlipaypfmOrderSyncAPIResponse 将 AlibabaAlihealthAlipaypfmOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAlipaypfmOrderSyncAPIResponse(v *AlibabaAlihealthAlipaypfmOrderSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAlipaypfmOrderSyncAPIResponse.Put(v)
}
