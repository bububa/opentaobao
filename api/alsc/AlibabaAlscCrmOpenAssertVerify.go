package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenAssertVerify 资产核销接口
// alibaba.alsc.crm.open.assert.verify
//
// 核销储值，积分，券资产
func AlibabaAlscCrmOpenAssertVerify(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenAssertVerifyAPIRequest, resp *alsc.AlibabaAlscCrmOpenAssertVerifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
