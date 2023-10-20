package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelorderupdatestatus 翱象商家自有渠道 订单状态更新
// alibaba.tcls.aelophy.merchant.channel.order.updatestatus
//
// 订单状态变更
func Alibabatclsaelophymerchantchannelorderupdatestatus(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
