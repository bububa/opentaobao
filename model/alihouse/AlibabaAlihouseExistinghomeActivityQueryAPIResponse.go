package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeactivityqueryAPIResponse 五福活动经纪人获取 API返回值
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
type AlibabaalihouseexistinghomeactivityqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeactivityqueryAPIResponseModel
}

// AlibabaalihouseexistinghomeactivityqueryAPIResponseModel is 五福活动经纪人获取 成功返回结果
type AlibabaalihouseexistinghomeactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihouseexistinghomeactivityqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
