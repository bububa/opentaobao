package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelOrderCreateAPIRequest
渠道分销创建订单接口 API请求
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供） */
type AlitripXhotelChannelOrderCreateAPIRequest struct {
	model.Params
	// 创建订单参数
	_outSourceOrderCreateReq *OutSourceOrderCreateReq
}

// New
