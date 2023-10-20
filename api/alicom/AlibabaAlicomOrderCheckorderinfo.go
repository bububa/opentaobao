package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomOrderCheckorderinfo 金融购机订单信息校验
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
func AlibabaAlicomOrderCheckorderinfo(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderCheckorderinfoAPIRequest, resp *alicom.AlibabaAlicomOrderCheckorderinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
