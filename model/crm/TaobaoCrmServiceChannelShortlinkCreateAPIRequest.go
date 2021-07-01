package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmServiceChannelShortlinkCreateAPIRequest
ECRM创建淘短链服务 API请求
taobao.crm.service.channel.shortlink.create

可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。 */
type TaobaoCrmServiceChannelShortlinkCreateAPIRequest struct {
	model.Params
	// 淘短链名称（最多只能16个中文字符，类型为订单链接时传入订单ID）。
	_shortLinkName string
	// 淘短链类型：LT_ITEM（商品淘短链）,LT_SHOP（店铺首页淘短链）,LT_ACTIVITY（活动页淘短链）,LT_TRADE（订单详情页淘短链）。
	_linkType string
	// 类型为LT_ITEM时必须传入商品ID，类型为LT_SHOP时必须传入空值，类型为LT_ACTIVITY时传入活动页URL（URL必须是taobao.com,tmall.com,jaeapp.com这三个域名下的URL），类型为LT_TRADE时传入订单ID。
	_shortLinkData string
}

// New
