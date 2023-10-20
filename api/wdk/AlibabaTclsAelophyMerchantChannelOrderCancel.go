package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelordercancel 翱象商家自有渠道 交易订单取消
// alibaba.tcls.aelophy.merchant.channel.order.cancel
//
// 翱象小程序用户取消订单
func Alibabatclsaelophymerchantchannelordercancel(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelordercancelAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelordercancelAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
