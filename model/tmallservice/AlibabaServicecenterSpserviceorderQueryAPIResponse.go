package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterspserviceorderqueryAPIResponse 服务供应链服务单查询 API返回值
// alibaba.servicecenter.spserviceorder.query
//
// 服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
type AlibabaservicecenterspserviceorderqueryAPIResponse struct {
	model.CommonResponse
	AlibabaservicecenterspserviceorderqueryAPIResponseModel
}

// AlibabaservicecenterspserviceorderqueryAPIResponseModel is 服务供应链服务单查询 成功返回结果
type AlibabaservicecenterspserviceorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_spserviceorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *AlibabaservicecenterspserviceorderqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
