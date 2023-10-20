package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelorderrefundconfirm 退款确认
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
func Alibabawdkchannelorderrefundconfirm(clt *core.SDKClient, req *wdk.AlibabawdkchannelorderrefundconfirmAPIRequest, session string) (*wdk.AlibabawdkchannelorderrefundconfirmAPIResponse, error) {
	var resp wdk.AlibabawdkchannelorderrefundconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
