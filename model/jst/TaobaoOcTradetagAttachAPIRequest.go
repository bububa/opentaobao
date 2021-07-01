package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcTradetagAttachAPIRequest
订单打标或者订单标签更新 API请求
taobao.oc.tradetag.attach

对订单添加标签和更新标签。标签分为官方标签和自定义标签。<br/>官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731<br/>自定义标签有2个通用属性:<br/>    `show_str:给消费者显示的字符串（如果可以显示的话）<br/>    `pic_urls:图片url,地址必须是图片空间的url,最多5张 */
type TaobaoOcTradetagAttachAPIRequest struct {
	model.Params
	// 标签类型       1：官方标签      2：自定义标签
	_tagType int64
	// 订单id
	_tid int64
	// 标签名称
	_tagName string
	// 标签值，json格式
	_tagValue string
	// 该标签在消费者端是否显示,0:不显示,1：显示
	_visible int64
}

// New
