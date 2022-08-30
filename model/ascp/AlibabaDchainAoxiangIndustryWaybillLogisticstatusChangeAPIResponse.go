package ascp

import (
	"encoding/xml"

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

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel is 物流节点回传 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_logisticstatus_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderConfirmResponse *TopResponse `json:"tms_order_confirm_response,omitempty" xml:"tms_order_confirm_response,omitempty"`
}
