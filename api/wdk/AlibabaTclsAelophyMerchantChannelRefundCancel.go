package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelrefundcancel 翱象商家自有渠道 逆向单申请取消
// alibaba.tcls.aelophy.merchant.channel.refund.cancel
//
// 翱象小程序 用户逆向申请取消
func Alibabatclsaelophymerchantchannelrefundcancel(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelrefundcancelAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelrefundcancelAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelrefundcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
