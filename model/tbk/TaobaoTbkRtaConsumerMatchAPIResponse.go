package tbk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTbkRtaConsumerMatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkRtaConsumerMatchAPIResponseModel).Reset()
}

// TaobaoTbkRtaConsumerMatchAPIResponseModel is 淘宝客-推广者-定向活动目标发布 成功返回结果
type TaobaoTbkRtaConsumerMatchAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_rta_consumer_match_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaoTbkRtaConsumerMatchData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkRtaConsumerMatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkRtaConsumerMatchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkRtaConsumerMatchAPIResponse)
	},
}

// GetTaobaoTbkRtaConsumerMatchAPIResponse 从 sync.Pool 获取 TaobaoTbkRtaConsumerMatchAPIResponse
func GetTaobaoTbkRtaConsumerMatchAPIResponse() *TaobaoTbkRtaConsumerMatchAPIResponse {
	return poolTaobaoTbkRtaConsumerMatchAPIResponse.Get().(*TaobaoTbkRtaConsumerMatchAPIResponse)
}

// ReleaseTaobaoTbkRtaConsumerMatchAPIResponse 将 TaobaoTbkRtaConsumerMatchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkRtaConsumerMatchAPIResponse(v *TaobaoTbkRtaConsumerMatchAPIResponse) {
	v.Reset()
	poolTaobaoTbkRtaConsumerMatchAPIResponse.Put(v)
}
