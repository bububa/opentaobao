package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixeticketdistributionqueryAPIResponse 分销电子票查询接口 API返回值
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
type AlibabadamaimaitixeticketdistributionqueryAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixeticketdistributionqueryAPIResponseModel
}

// AlibabadamaimaitixeticketdistributionqueryAPIResponseModel is 分销电子票查询接口 成功返回结果
type AlibabadamaimaitixeticketdistributionqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_eticket_distribution_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
