package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantchannelordersliceget 获取运力时间片信息
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
func Alibabatclsaelophymerchantchannelordersliceget(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantchannelorderslicegetAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantchannelorderslicegetAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantchannelorderslicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
