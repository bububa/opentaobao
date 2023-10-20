package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrywaybillpreacceptAPIResponse 商家ERP预推单 API返回值
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
type AlibabaascpindustrywaybillpreacceptAPIResponse struct {
	model.CommonResponse
	AlibabaascpindustrywaybillpreacceptAPIResponseModel
}

// AlibabaascpindustrywaybillpreacceptAPIResponseModel is 商家ERP预推单 成功返回结果
type AlibabaascpindustrywaybillpreacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_waybill_pre_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
