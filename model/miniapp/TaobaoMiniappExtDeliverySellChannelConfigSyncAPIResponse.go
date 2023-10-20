package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse 写入商家配置信息 API返回值
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
type TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponseModel).Reset()
}

// TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponseModel is 写入商家配置信息 成功返回结果
type TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_sell_channel_config_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
	// 操作结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorType = 0
	m.Successful = false
	m.Model = false
}

var poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse)
	},
}

// GetTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse
func GetTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse() *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse {
	return poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse.Get().(*TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse)
}

// ReleaseTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse 将 TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse(v *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse.Put(v)
}
