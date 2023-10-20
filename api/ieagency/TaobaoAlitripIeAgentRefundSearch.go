package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundsearch 卖家查询退票申请
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
func Taobaoalitripieagentrefundsearch(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundsearchAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundsearchAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
