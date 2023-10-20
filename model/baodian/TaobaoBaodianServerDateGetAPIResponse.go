package baodian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianServerDateGetAPIResponse 服务器时间获取 API返回值
// taobao.baodian.server.date.get
//
// 获取服务器时间
type TaobaoBaodianServerDateGetAPIResponse struct {
	model.CommonResponse
	TaobaoBaodianServerDateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaodianServerDateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaodianServerDateGetAPIResponseModel).Reset()
}

// TaobaoBaodianServerDateGetAPIResponseModel is 服务器时间获取 成功返回结果
type TaobaoBaodianServerDateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"baodian_server_date_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回时间为毫秒
	ServerDate int64 `json:"server_date,omitempty" xml:"server_date,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaodianServerDateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServerDate = 0
}

var poolTaobaoBaodianServerDateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaodianServerDateGetAPIResponse)
	},
}

// GetTaobaoBaodianServerDateGetAPIResponse 从 sync.Pool 获取 TaobaoBaodianServerDateGetAPIResponse
func GetTaobaoBaodianServerDateGetAPIResponse() *TaobaoBaodianServerDateGetAPIResponse {
	return poolTaobaoBaodianServerDateGetAPIResponse.Get().(*TaobaoBaodianServerDateGetAPIResponse)
}

// ReleaseTaobaoBaodianServerDateGetAPIResponse 将 TaobaoBaodianServerDateGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaodianServerDateGetAPIResponse(v *TaobaoBaodianServerDateGetAPIResponse) {
	v.Reset()
	poolTaobaoBaodianServerDateGetAPIResponse.Put(v)
}
