package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkRtaConsumerMatchAPIResponse 淘宝客-推广者-定向活动目标发布 API返回值
// taobao.tbk.rta.consumer.match
//
// 淘客计划向用户推送某个定向活动时，调用该接口判断用户是否符合活动目标（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkRtaConsumerMatchAPIResponse struct {
	model.CommonResponse
	TaobaoTbkRtaConsumerMatchAPIResponseModel
}

// TaobaoTbkRtaConsumerMatchAPIResponseModel is 淘宝客-推广者-定向活动目标发布 成功返回结果
type TaobaoTbkRtaConsumerMatchAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_rta_consumer_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaoTbkRtaConsumerMatchData `json:"data,omitempty" xml:"data,omitempty"`
}
