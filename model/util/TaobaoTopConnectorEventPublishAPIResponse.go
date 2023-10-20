package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopConnectorEventPublishAPIResponse 连接器事件发布 API返回值
// taobao.top.connector.event.publish
//
// 连接器事件发布
type TaobaoTopConnectorEventPublishAPIResponse struct {
	model.CommonResponse
	TaobaoTopConnectorEventPublishAPIResponseModel
}

// TaobaoTopConnectorEventPublishAPIResponseModel is 连接器事件发布 成功返回结果
type TaobaoTopConnectorEventPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"top_connector_event_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发布事件响应
	Result *EventPublishThirdPartyResponse `json:"result,omitempty" xml:"result,omitempty"`
}
