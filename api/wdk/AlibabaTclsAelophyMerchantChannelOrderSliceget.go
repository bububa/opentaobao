package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantChannelOrderSliceget 获取运力时间片信息
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
func AlibabaTclsAelophyMerchantChannelOrderSliceget(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIRequest, session string) (*wdk.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
