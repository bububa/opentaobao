package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmUpdateAlipayCarno 更新支付宝业务卡号
// alibaba.westcrm.update.alipay.carno
//
// 更新支付宝业务卡号
func AlibabaWestcrmUpdateAlipayCarno(clt *core.SDKClient, req *westcrm.AlibabaWestcrmUpdateAlipayCarnoAPIRequest, session string) (*westcrm.AlibabaWestcrmUpdateAlipayCarnoAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmUpdateAlipayCarnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
