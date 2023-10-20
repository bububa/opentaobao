package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Alitripxhotelchannelnotify 分销渠道各类通知接口
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
func Alitripxhotelchannelnotify(clt *core.SDKClient, req *xhotelonlineorder.AlitripxhotelchannelnotifyAPIRequest, session string) (*xhotelonlineorder.AlitripxhotelchannelnotifyAPIResponse, error) {
	var resp xhotelonlineorder.AlitripxhotelchannelnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
