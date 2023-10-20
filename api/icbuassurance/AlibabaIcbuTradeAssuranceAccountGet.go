package icbuassurance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuassurance"
)

// Alibabaicbutradeassuranceaccountget icbu信保账户信息
// alibaba.icbu.trade.assurance.account.get
//
// icbu交易信用保障开通状态&amp;额度信息查询
func Alibabaicbutradeassuranceaccountget(clt *core.SDKClient, req *icbuassurance.AlibabaicbutradeassuranceaccountgetAPIRequest, session string) (*icbuassurance.AlibabaicbutradeassuranceaccountgetAPIResponse, error) {
	var resp icbuassurance.AlibabaicbutradeassuranceaccountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
