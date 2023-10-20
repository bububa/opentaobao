package larkiot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderGetcinemasAPIResponse 获取iot渠道开放的影院 API返回值
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
type TaobaoLarkIotOrderGetcinemasAPIResponse struct {
	model.CommonResponse
	TaobaoLarkIotOrderGetcinemasAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLarkIotOrderGetcinemasAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLarkIotOrderGetcinemasAPIResponseModel).Reset()
}

// TaobaoLarkIotOrderGetcinemasAPIResponseModel is 获取iot渠道开放的影院 成功返回结果
type TaobaoLarkIotOrderGetcinemasAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_iot_order_getcinemas_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 有权限的影院列表
	Data *BizListResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLarkIotOrderGetcinemasAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoLarkIotOrderGetcinemasAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLarkIotOrderGetcinemasAPIResponse)
	},
}

// GetTaobaoLarkIotOrderGetcinemasAPIResponse 从 sync.Pool 获取 TaobaoLarkIotOrderGetcinemasAPIResponse
func GetTaobaoLarkIotOrderGetcinemasAPIResponse() *TaobaoLarkIotOrderGetcinemasAPIResponse {
	return poolTaobaoLarkIotOrderGetcinemasAPIResponse.Get().(*TaobaoLarkIotOrderGetcinemasAPIResponse)
}

// ReleaseTaobaoLarkIotOrderGetcinemasAPIResponse 将 TaobaoLarkIotOrderGetcinemasAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLarkIotOrderGetcinemasAPIResponse(v *TaobaoLarkIotOrderGetcinemasAPIResponse) {
	v.Reset()
	poolTaobaoLarkIotOrderGetcinemasAPIResponse.Put(v)
}
