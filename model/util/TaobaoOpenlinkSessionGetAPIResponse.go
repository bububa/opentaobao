package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenlinkSessionGetAPIResponse 获取授权session信息 API返回值
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
type TaobaoOpenlinkSessionGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenlinkSessionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenlinkSessionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenlinkSessionGetAPIResponseModel).Reset()
}

// TaobaoOpenlinkSessionGetAPIResponseModel is 获取授权session信息 成功返回结果
type TaobaoOpenlinkSessionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openlink_session_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOpenlinkSessionGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenlinkSessionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenlinkSessionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenlinkSessionGetAPIResponse)
	},
}

// GetTaobaoOpenlinkSessionGetAPIResponse 从 sync.Pool 获取 TaobaoOpenlinkSessionGetAPIResponse
func GetTaobaoOpenlinkSessionGetAPIResponse() *TaobaoOpenlinkSessionGetAPIResponse {
	return poolTaobaoOpenlinkSessionGetAPIResponse.Get().(*TaobaoOpenlinkSessionGetAPIResponse)
}

// ReleaseTaobaoOpenlinkSessionGetAPIResponse 将 TaobaoOpenlinkSessionGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenlinkSessionGetAPIResponse(v *TaobaoOpenlinkSessionGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenlinkSessionGetAPIResponse.Put(v)
}
