package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopEventPublishAPIResponse 同步事件发布 API返回值
// taobao.top.event.publish
//
// 同步事件发布
type TaobaoTopEventPublishAPIResponse struct {
	model.CommonResponse
	TaobaoTopEventPublishAPIResponseModel
}

// TaobaoTopEventPublishAPIResponseModel is 同步事件发布 成功返回结果
type TaobaoTopEventPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"top_event_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 事件返回值
	Data *EventPublishResponse `json:"data,omitempty" xml:"data,omitempty"`
}
