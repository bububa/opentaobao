package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillEditAPIResponse 服务商编辑运单 API返回值
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
type AlibabaDchainAoxiangIndustryWaybillEditAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangIndustryWaybillEditAPIResponseModel
}

// AlibabaDchainAoxiangIndustryWaybillEditAPIResponseModel is 服务商编辑运单 成功返回结果
type AlibabaDchainAoxiangIndustryWaybillEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_industry_waybill_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	TmsOrderEditResponse *TopResponse `json:"tms_order_edit_response,omitempty" xml:"tms_order_edit_response,omitempty"`
}
