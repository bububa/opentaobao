package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrywaybillacceptAPIResponse 商家ERP预推单 API返回值
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
type AlibabaascpindustrywaybillacceptAPIResponse struct {
	model.CommonResponse
	AlibabaascpindustrywaybillacceptAPIResponseModel
}

// AlibabaascpindustrywaybillacceptAPIResponseModel is 商家ERP预推单 成功返回结果
type AlibabaascpindustrywaybillacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_waybill_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
