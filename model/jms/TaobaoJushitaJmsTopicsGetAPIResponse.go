package jms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsTopicsGetAPIResponse 根据用户nick获取开通的消息列表 API返回值
// taobao.jushita.jms.topics.get
//
// 根据用户nick获取开通的消息列表
type TaobaoJushitaJmsTopicsGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsTopicsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsTopicsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJmsTopicsGetAPIResponseModel).Reset()
}

// TaobaoJushitaJmsTopicsGetAPIResponseModel is 根据用户nick获取开通的消息列表 成功返回结果
type TaobaoJushitaJmsTopicsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_topics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// topic列表
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// 错误信息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsTopicsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.ResultMessage = ""
	m.ResultCode = ""
}

var poolTaobaoJushitaJmsTopicsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJmsTopicsGetAPIResponse)
	},
}

// GetTaobaoJushitaJmsTopicsGetAPIResponse 从 sync.Pool 获取 TaobaoJushitaJmsTopicsGetAPIResponse
func GetTaobaoJushitaJmsTopicsGetAPIResponse() *TaobaoJushitaJmsTopicsGetAPIResponse {
	return poolTaobaoJushitaJmsTopicsGetAPIResponse.Get().(*TaobaoJushitaJmsTopicsGetAPIResponse)
}

// ReleaseTaobaoJushitaJmsTopicsGetAPIResponse 将 TaobaoJushitaJmsTopicsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJmsTopicsGetAPIResponse(v *TaobaoJushitaJmsTopicsGetAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJmsTopicsGetAPIResponse.Put(v)
}
