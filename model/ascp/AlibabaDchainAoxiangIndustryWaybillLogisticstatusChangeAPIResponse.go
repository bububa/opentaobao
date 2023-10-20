package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse 物流节点回传 API返回值
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel is 物流节点回传 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_logisticstatus_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderConfirmResponse *TopResponse `json:"tms_order_confirm_response,omitempty" xml:"tms_order_confirm_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmsOrderConfirmResponse = nil
}

var poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse)
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse
func GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse {
	return poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse.Get().(*AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse 将 AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse(v *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse.Put(v)
}
