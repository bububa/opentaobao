package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsCpGetAPIResponse 快递公司资源列表查询接口 API返回值
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
type AlibabaAscpLogisticsCpGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsCpGetAPIResponseModel
}

// AlibabaAscpLogisticsCpGetAPIResponseModel is 快递公司资源列表查询接口 成功返回结果
type AlibabaAscpLogisticsCpGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_cp_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
