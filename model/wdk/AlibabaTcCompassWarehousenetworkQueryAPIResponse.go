package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatccompasswarehousenetworkqueryAPIResponse 按仓维度来查询鸟潮网络 API返回值
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
type AlibabatccompasswarehousenetworkqueryAPIResponse struct {
	model.CommonResponse
	AlibabatccompasswarehousenetworkqueryAPIResponseModel
}

// AlibabatccompasswarehousenetworkqueryAPIResponseModel is 按仓维度来查询鸟潮网络 成功返回结果
type AlibabatccompasswarehousenetworkqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tc_compass_warehousenetwork_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
