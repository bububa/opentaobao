package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmServiceChannelShortlinkCreateAPIResponse
ECRM创建淘短链服务 API返回值
taobao.crm.service.channel.shortlink.create

可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。 */
type TaobaoCrmServiceChannelShortlinkCreateAPIResponse struct {
	model.CommonResponse
	TaobaoCrmServiceChannelShortlinkCreateAPIResponseModel
}

// TaobaoCrmServiceChannelShortlinkCreateAPIResponseModel is ECRM创建淘短链服务 成功返回结果
type TaobaoCrmServiceChannelShortlinkCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_service_channel_shortlink_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的淘短链。
	ShortLink string `json:"short_link,omitempty" xml:"short_link,omitempty"`
}
