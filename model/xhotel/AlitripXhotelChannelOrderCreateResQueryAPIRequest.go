package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelOrderCreateResQueryAPIRequest
分销订单查询订单创建结果 API请求
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。 */
type AlitripXhotelChannelOrderCreateResQueryAPIRequest struct {
	model.Params
	// 外部渠道订单号
	_outSourceOrderId string
}

// New
