package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse ISV写入应用的渠道信息 API返回值
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
type TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappExtDeliveryAppChannelSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappExtDeliveryAppChannelSyncAPIResponseModel).Reset()
}

// TaobaoMiniappExtDeliveryAppChannelSyncAPIResponseModel is ISV写入应用的渠道信息 成功返回结果
type TaobaoMiniappExtDeliveryAppChannelSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_app_channel_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// configId
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliveryAppChannelSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorType = 0
	m.Model = 0
	m.Successful = false
}

var poolTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse)
	},
}

// GetTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse 从 sync.Pool 获取 TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse
func GetTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse() *TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse {
	return poolTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse.Get().(*TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse)
}

// ReleaseTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse 将 TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse(v *TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse) {
	v.Reset()
	poolTaobaoMiniappExtDeliveryAppChannelSyncAPIResponse.Put(v)
}
