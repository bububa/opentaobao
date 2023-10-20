package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsNotraceSendAPIResponse 供应商-异云-无需物流发货 API返回值
// alibaba.lst.logistics.notrace.send
//
// 异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
type AlibabaLstLogisticsNotraceSendAPIResponse struct {
	model.CommonResponse
	AlibabaLstLogisticsNotraceSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsNotraceSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstLogisticsNotraceSendAPIResponseModel).Reset()
}

// AlibabaLstLogisticsNotraceSendAPIResponseModel is 供应商-异云-无需物流发货 成功返回结果
type AlibabaLstLogisticsNotraceSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_notrace_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstLogisticsNotraceSendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsNotraceSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstLogisticsNotraceSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsNotraceSendAPIResponse)
	},
}

// GetAlibabaLstLogisticsNotraceSendAPIResponse 从 sync.Pool 获取 AlibabaLstLogisticsNotraceSendAPIResponse
func GetAlibabaLstLogisticsNotraceSendAPIResponse() *AlibabaLstLogisticsNotraceSendAPIResponse {
	return poolAlibabaLstLogisticsNotraceSendAPIResponse.Get().(*AlibabaLstLogisticsNotraceSendAPIResponse)
}

// ReleaseAlibabaLstLogisticsNotraceSendAPIResponse 将 AlibabaLstLogisticsNotraceSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstLogisticsNotraceSendAPIResponse(v *AlibabaLstLogisticsNotraceSendAPIResponse) {
	v.Reset()
	poolAlibabaLstLogisticsNotraceSendAPIResponse.Put(v)
}
