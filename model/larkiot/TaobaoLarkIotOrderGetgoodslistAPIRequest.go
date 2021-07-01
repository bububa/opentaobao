package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkIotOrderGetgoodslistAPIRequest
iot渠道获取卖品信息 API请求
taobao.lark.iot.order.getgoodslist

iot无人超市服务商通过接口获取影院的可售卖品数据 */
type TaobaoLarkIotOrderGetgoodslistAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
	// 影院内码
	_cinemaLinkId string
}

// New
