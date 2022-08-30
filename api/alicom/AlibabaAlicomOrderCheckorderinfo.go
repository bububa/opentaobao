package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomOrderCheckorderinfo 金融购机订单信息校验
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
func AlibabaAlicomOrderCheckorderinfo(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderCheckorderinfoAPIRequest, session string) (*alicom.AlibabaAlicomOrderCheckorderinfoAPIResponse, error) {
	var resp alicom.AlibabaAlicomOrderCheckorderinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
