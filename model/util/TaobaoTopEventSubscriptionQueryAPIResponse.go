package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopeventsubscriptionqueryAPIResponse 三方事件订阅查询 API返回值
// taobao.top.event.subscription.query
//
// 三方事件订阅查询
type TaobaotopeventsubscriptionqueryAPIResponse struct {
	model.CommonResponse
	TaobaotopeventsubscriptionqueryAPIResponseModel
}

// TaobaotopeventsubscriptionqueryAPIResponseModel is 三方事件订阅查询 成功返回结果
type TaobaotopeventsubscriptionqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_event_subscription_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否订阅服务
	StartService bool `json:"start_service,omitempty" xml:"start_service,omitempty"`
}
