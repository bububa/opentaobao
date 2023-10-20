package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse 查询商家配置的信息 API返回值
// taobao.miniapp.ext.delivery.sell.channel.configs.query
//
// 查询商家配置的信息
type TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponseModel).Reset()
}

// TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponseModel is 查询商家配置的信息 成功返回结果
type TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_sell_channel_configs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MiniappResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse)
	},
}

// GetTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse
func GetTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse() *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse {
	return poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse.Get().(*TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse)
}

// ReleaseTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse 将 TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse(v *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse.Put(v)
}
