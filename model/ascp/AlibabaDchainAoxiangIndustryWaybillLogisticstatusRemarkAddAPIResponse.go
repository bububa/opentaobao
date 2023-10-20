package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse 客服增加备注 API返回值
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel is 客服增加备注 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_logisticstatus_remark_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderRemarkResponse *TopResponse `json:"tms_order_remark_response,omitempty" xml:"tms_order_remark_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmsOrderRemarkResponse = nil
}

var poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse)
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse
func GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse {
	return poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse.Get().(*AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse 将 AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse(v *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse.Put(v)
}
