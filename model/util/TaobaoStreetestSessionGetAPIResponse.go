package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoStreetestSessionGetAPIResponse 根据获取压测用户的sessionKey API返回值
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaoStreetestSessionGetAPIResponse struct {
	model.CommonResponse
	TaobaoStreetestSessionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoStreetestSessionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoStreetestSessionGetAPIResponseModel).Reset()
}

// TaobaoStreetestSessionGetAPIResponseModel is 根据获取压测用户的sessionKey 成功返回结果
type TaobaoStreetestSessionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"streetest_session_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 压测账号对应的sessionKey
	StreeTestSessionKey string `json:"stree_test_session_key,omitempty" xml:"stree_test_session_key,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoStreetestSessionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StreeTestSessionKey = ""
}

var poolTaobaoStreetestSessionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoStreetestSessionGetAPIResponse)
	},
}

// GetTaobaoStreetestSessionGetAPIResponse 从 sync.Pool 获取 TaobaoStreetestSessionGetAPIResponse
func GetTaobaoStreetestSessionGetAPIResponse() *TaobaoStreetestSessionGetAPIResponse {
	return poolTaobaoStreetestSessionGetAPIResponse.Get().(*TaobaoStreetestSessionGetAPIResponse)
}

// ReleaseTaobaoStreetestSessionGetAPIResponse 将 TaobaoStreetestSessionGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoStreetestSessionGetAPIResponse(v *TaobaoStreetestSessionGetAPIResponse) {
	v.Reset()
	poolTaobaoStreetestSessionGetAPIResponse.Put(v)
}
