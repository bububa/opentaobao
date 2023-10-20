package ascp

import (
	"encoding/xml"

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

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel is 客服增加备注 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_logisticstatus_remark_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderRemarkResponse *TopResponse `json:"tms_order_remark_response,omitempty" xml:"tms_order_remark_response,omitempty"`
}
