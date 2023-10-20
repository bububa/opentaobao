package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWaybillAcceptAPIResponse 商家ERP预推单 API返回值
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
type AlibabaAscpIndustryWaybillAcceptAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryWaybillAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryWaybillAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryWaybillAcceptAPIResponseModel).Reset()
}

// AlibabaAscpIndustryWaybillAcceptAPIResponseModel is 商家ERP预推单 成功返回结果
type AlibabaAscpIndustryWaybillAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_waybill_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryWaybillAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryWaybillAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryWaybillAcceptAPIResponse)
	},
}

// GetAlibabaAscpIndustryWaybillAcceptAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryWaybillAcceptAPIResponse
func GetAlibabaAscpIndustryWaybillAcceptAPIResponse() *AlibabaAscpIndustryWaybillAcceptAPIResponse {
	return poolAlibabaAscpIndustryWaybillAcceptAPIResponse.Get().(*AlibabaAscpIndustryWaybillAcceptAPIResponse)
}

// ReleaseAlibabaAscpIndustryWaybillAcceptAPIResponse 将 AlibabaAscpIndustryWaybillAcceptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryWaybillAcceptAPIResponse(v *AlibabaAscpIndustryWaybillAcceptAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryWaybillAcceptAPIResponse.Put(v)
}
