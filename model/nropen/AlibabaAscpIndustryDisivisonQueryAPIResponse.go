package nropen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryDisivisonQueryAPIResponse 查询服务支持地区列表 API返回值
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
type AlibabaAscpIndustryDisivisonQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryDisivisonQueryAPIResponseModel
}

// AlibabaAscpIndustryDisivisonQueryAPIResponseModel is 查询服务支持地区列表 成功返回结果
type AlibabaAscpIndustryDisivisonQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_disivison_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *AlibabaAscpIndustryDisivisonQueryData `json:"data,omitempty" xml:"data,omitempty"`
}
