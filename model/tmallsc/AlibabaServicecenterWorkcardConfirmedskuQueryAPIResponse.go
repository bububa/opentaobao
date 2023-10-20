package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse 查询确认履行的服务项 API返回值
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
type AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel
}

// AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel is 查询确认履行的服务项 成功返回结果
type AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_confirmedsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterWorkcardConfirmedskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
