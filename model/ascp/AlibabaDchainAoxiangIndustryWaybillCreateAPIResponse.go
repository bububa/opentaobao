package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse 服务商开运单 API返回值
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
type AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangIndustryWaybillCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangIndustryWaybillCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangIndustryWaybillCreateAPIResponseModel is 服务商开运单 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderCreateResponse *TopResponse `json:"tms_order_create_response,omitempty" xml:"tms_order_create_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmsOrderCreateResponse = nil
}

var poolAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse
func GetAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse() *AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse {
	return poolAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse.Get().(*AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse 将 AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse(v *AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillCreateAPIResponse.Put(v)
}
