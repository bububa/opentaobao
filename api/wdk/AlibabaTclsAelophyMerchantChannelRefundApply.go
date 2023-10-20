package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelrefundapply 翱象商家自有渠道 逆向单申请
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
func Alibabatclsaelophymerchantchannelrefundapply(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelrefundapplyAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelrefundapplyAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelrefundapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
