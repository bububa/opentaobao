package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOrderurlGetAPIResponse 百川订单详情 API返回值
// taobao.baichuan.orderurl.get
//
// 百川订单详情
type TaobaoBaichuanOrderurlGetAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOrderurlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOrderurlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOrderurlGetAPIResponseModel).Reset()
}

// TaobaoBaichuanOrderurlGetAPIResponseModel is 百川订单详情 成功返回结果
type TaobaoBaichuanOrderurlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_orderurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOrderurlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOrderurlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOrderurlGetAPIResponse)
	},
}

// GetTaobaoBaichuanOrderurlGetAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOrderurlGetAPIResponse
func GetTaobaoBaichuanOrderurlGetAPIResponse() *TaobaoBaichuanOrderurlGetAPIResponse {
	return poolTaobaoBaichuanOrderurlGetAPIResponse.Get().(*TaobaoBaichuanOrderurlGetAPIResponse)
}

// ReleaseTaobaoBaichuanOrderurlGetAPIResponse 将 TaobaoBaichuanOrderurlGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOrderurlGetAPIResponse(v *TaobaoBaichuanOrderurlGetAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOrderurlGetAPIResponse.Put(v)
}
