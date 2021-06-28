package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ECRM创建淘短链服务 APIResponse
taobao.crm.service.channel.shortlink.create

可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。
*/
type TaobaoCrmServiceChannelShortlinkCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmServiceChannelShortlinkCreateResponse `json:"crm_service_channel_shortlink_create_response,omitempty"` 
    TaobaoCrmServiceChannelShortlinkCreateResponse
}

/* model for simplify = false
type TaobaoCrmServiceChannelShortlinkCreateResponse struct {

    // 返回的淘短链。
    
    ShortLink   string `json:"short_link,omitempty"`
    

}
*/

type TaobaoCrmServiceChannelShortlinkCreateResponse struct {

    // 返回的淘短链。
    ShortLink   string `json:"short_link,omitempty"`

}
