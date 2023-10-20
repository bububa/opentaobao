package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenAssertRefund 资产核销回退接口
// alibaba.alsc.crm.open.assert.refund
//
// 回退已经核销的储值积分券资产
func AlibabaAlscCrmOpenAssertRefund(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenAssertRefundAPIRequest, resp *alsc.AlibabaAlscCrmOpenAssertRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
