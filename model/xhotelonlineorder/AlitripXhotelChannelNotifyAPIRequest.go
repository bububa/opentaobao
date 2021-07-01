package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelNotifyAPIRequest
分销渠道各类通知接口 API请求
alitrip.xhotel.channel.notify

分销渠道支付通知 */
type AlitripXhotelChannelNotifyAPIRequest struct {
	model.Params
	// 通知类型查询条件
	_orderNotifyQuery *OrderNotifyQuery
}

// New
