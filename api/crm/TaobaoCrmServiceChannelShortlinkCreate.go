package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmServiceChannelShortlinkCreate ECRM创建淘短链服务
// taobao.crm.service.channel.shortlink.create
//
// 可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。
func TaobaoCrmServiceChannelShortlinkCreate(clt *core.SDKClient, req *crm.TaobaoCrmServiceChannelShortlinkCreateAPIRequest, resp *crm.TaobaoCrmServiceChannelShortlinkCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
