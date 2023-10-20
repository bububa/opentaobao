package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousebusinessactivityqueryAPIResponse 电商券活动公司数据查询 API返回值
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
type AlibabaalihousebusinessactivityqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihousebusinessactivityqueryAPIResponseModel
}

// AlibabaalihousebusinessactivityqueryAPIResponseModel is 电商券活动公司数据查询 成功返回结果
type AlibabaalihousebusinessactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_business_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousebusinessactivityqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
