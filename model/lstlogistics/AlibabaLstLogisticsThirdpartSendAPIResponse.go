package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsThirdpartSendAPIResponse 供应商-异云-使用第三方物流发货 API返回值
// alibaba.lst.logistics.thirdpart.send
//
// 异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
type AlibabaLstLogisticsThirdpartSendAPIResponse struct {
	model.CommonResponse
	AlibabaLstLogisticsThirdpartSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsThirdpartSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstLogisticsThirdpartSendAPIResponseModel).Reset()
}

// AlibabaLstLogisticsThirdpartSendAPIResponseModel is 供应商-异云-使用第三方物流发货 成功返回结果
type AlibabaLstLogisticsThirdpartSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_thirdpart_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstLogisticsThirdpartSendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsThirdpartSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstLogisticsThirdpartSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsThirdpartSendAPIResponse)
	},
}

// GetAlibabaLstLogisticsThirdpartSendAPIResponse 从 sync.Pool 获取 AlibabaLstLogisticsThirdpartSendAPIResponse
func GetAlibabaLstLogisticsThirdpartSendAPIResponse() *AlibabaLstLogisticsThirdpartSendAPIResponse {
	return poolAlibabaLstLogisticsThirdpartSendAPIResponse.Get().(*AlibabaLstLogisticsThirdpartSendAPIResponse)
}

// ReleaseAlibabaLstLogisticsThirdpartSendAPIResponse 将 AlibabaLstLogisticsThirdpartSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstLogisticsThirdpartSendAPIResponse(v *AlibabaLstLogisticsThirdpartSendAPIResponse) {
	v.Reset()
	poolAlibabaLstLogisticsThirdpartSendAPIResponse.Put(v)
}
