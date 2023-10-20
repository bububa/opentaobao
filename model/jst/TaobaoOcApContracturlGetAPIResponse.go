package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcApContracturlGetAPIResponse 按用户获取支付宝代扣协议链接地址 API返回值
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
type TaobaoOcApContracturlGetAPIResponse struct {
	model.CommonResponse
	TaobaoOcApContracturlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOcApContracturlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOcApContracturlGetAPIResponseModel).Reset()
}

// TaobaoOcApContracturlGetAPIResponseModel is 按用户获取支付宝代扣协议链接地址 成功返回结果
type TaobaoOcApContracturlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_ap_contracturl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述信息
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
	// 代扣url地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 判断操作是否执行成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOcApContracturlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDescription = ""
	m.Url = ""
	m.IsSuccess = false
}

var poolTaobaoOcApContracturlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOcApContracturlGetAPIResponse)
	},
}

// GetTaobaoOcApContracturlGetAPIResponse 从 sync.Pool 获取 TaobaoOcApContracturlGetAPIResponse
func GetTaobaoOcApContracturlGetAPIResponse() *TaobaoOcApContracturlGetAPIResponse {
	return poolTaobaoOcApContracturlGetAPIResponse.Get().(*TaobaoOcApContracturlGetAPIResponse)
}

// ReleaseTaobaoOcApContracturlGetAPIResponse 将 TaobaoOcApContracturlGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOcApContracturlGetAPIResponse(v *TaobaoOcApContracturlGetAPIResponse) {
	v.Reset()
	poolTaobaoOcApContracturlGetAPIResponse.Put(v)
}
