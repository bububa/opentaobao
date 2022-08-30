package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse 查询工单关联的服务项 API返回值
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
type AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel
}

// AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel is 查询工单关联的服务项 成功返回结果
type AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_relatedsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterWorkcardRelatedskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
