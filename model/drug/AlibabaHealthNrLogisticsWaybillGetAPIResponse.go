package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsWaybillGetAPIResponse 电子面单查询接口 API返回值
// alibaba.health.nr.logistics.waybill.get
//
// 商家登录后根据订单号查询物流单号及电子面单信息
type AlibabaHealthNrLogisticsWaybillGetAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrLogisticsWaybillGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsWaybillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrLogisticsWaybillGetAPIResponseModel).Reset()
}

// AlibabaHealthNrLogisticsWaybillGetAPIResponseModel is 电子面单查询接口 成功返回结果
type AlibabaHealthNrLogisticsWaybillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_logistics_waybill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果对象
	ResponseResult *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrLogisticsWaybillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseResult = nil
}

var poolAlibabaHealthNrLogisticsWaybillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrLogisticsWaybillGetAPIResponse)
	},
}

// GetAlibabaHealthNrLogisticsWaybillGetAPIResponse 从 sync.Pool 获取 AlibabaHealthNrLogisticsWaybillGetAPIResponse
func GetAlibabaHealthNrLogisticsWaybillGetAPIResponse() *AlibabaHealthNrLogisticsWaybillGetAPIResponse {
	return poolAlibabaHealthNrLogisticsWaybillGetAPIResponse.Get().(*AlibabaHealthNrLogisticsWaybillGetAPIResponse)
}

// ReleaseAlibabaHealthNrLogisticsWaybillGetAPIResponse 将 AlibabaHealthNrLogisticsWaybillGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrLogisticsWaybillGetAPIResponse(v *AlibabaHealthNrLogisticsWaybillGetAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrLogisticsWaybillGetAPIResponse.Put(v)
}
