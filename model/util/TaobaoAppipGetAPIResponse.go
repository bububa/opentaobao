package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppipGetAPIResponse 获取ISV发起请求服务器IP API返回值
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
type TaobaoAppipGetAPIResponse struct {
	model.CommonResponse
	TaobaoAppipGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAppipGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAppipGetAPIResponseModel).Reset()
}

// TaobaoAppipGetAPIResponseModel is 获取ISV发起请求服务器IP 成功返回结果
type TaobaoAppipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"appip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ISV发起请求服务器IP
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAppipGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ip = ""
}

var poolTaobaoAppipGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAppipGetAPIResponse)
	},
}

// GetTaobaoAppipGetAPIResponse 从 sync.Pool 获取 TaobaoAppipGetAPIResponse
func GetTaobaoAppipGetAPIResponse() *TaobaoAppipGetAPIResponse {
	return poolTaobaoAppipGetAPIResponse.Get().(*TaobaoAppipGetAPIResponse)
}

// ReleaseTaobaoAppipGetAPIResponse 将 TaobaoAppipGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAppipGetAPIResponse(v *TaobaoAppipGetAPIResponse) {
	v.Reset()
	poolTaobaoAppipGetAPIResponse.Put(v)
}
