package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWaybillPreAcceptAPIResponse 商家ERP预推单 API返回值
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
type AlibabaAscpIndustryWaybillPreAcceptAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryWaybillPreAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryWaybillPreAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryWaybillPreAcceptAPIResponseModel).Reset()
}

// AlibabaAscpIndustryWaybillPreAcceptAPIResponseModel is 商家ERP预推单 成功返回结果
type AlibabaAscpIndustryWaybillPreAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_waybill_pre_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryWaybillPreAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryWaybillPreAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryWaybillPreAcceptAPIResponse)
	},
}

// GetAlibabaAscpIndustryWaybillPreAcceptAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryWaybillPreAcceptAPIResponse
func GetAlibabaAscpIndustryWaybillPreAcceptAPIResponse() *AlibabaAscpIndustryWaybillPreAcceptAPIResponse {
	return poolAlibabaAscpIndustryWaybillPreAcceptAPIResponse.Get().(*AlibabaAscpIndustryWaybillPreAcceptAPIResponse)
}

// ReleaseAlibabaAscpIndustryWaybillPreAcceptAPIResponse 将 AlibabaAscpIndustryWaybillPreAcceptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryWaybillPreAcceptAPIResponse(v *AlibabaAscpIndustryWaybillPreAcceptAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryWaybillPreAcceptAPIResponse.Put(v)
}
