package miniapp

import (
	"encoding/xml"

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

// TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel is ISV查询应用的渠道信息 成功返回结果
type TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_app_channel_configs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// body
	Result *MiniappResult `json:"result,omitempty" xml:"result,omitempty"`
}
