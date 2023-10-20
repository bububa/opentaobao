package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse ISV查询应用的渠道信息 API返回值
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
type TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel).Reset()
}

// TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel is ISV查询应用的渠道信息 成功返回结果
type TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_app_channel_configs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// body
	Result *MiniappResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse)
	},
}

// GetTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse 从 sync.Pool 获取 TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse
func GetTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse() *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse {
	return poolTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse.Get().(*TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse)
}

// ReleaseTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse 将 TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse(v *TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse) {
	v.Reset()
	poolTaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse.Put(v)
}
