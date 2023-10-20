package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// AlitripXhotelChannelNotify 分销渠道各类通知接口
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
func AlitripXhotelChannelNotify(clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelNotifyAPIRequest, resp *xhotelonlineorder.AlitripXhotelChannelNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
