package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// AlitripXhotelChannelOrderCreate 渠道分销创建订单接口
// alitrip.xhotel.channel.order.create
//
// 创建订单接口服务（如菲住等其他渠道分销提供）
func AlitripXhotelChannelOrderCreate(clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelOrderCreateAPIRequest, resp *xhotelonlineorder.AlitripXhotelChannelOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
