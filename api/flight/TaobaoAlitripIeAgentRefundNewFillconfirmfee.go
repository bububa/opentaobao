package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaoalitripieagentrefundnewfillconfirmfee 新模型-回填申请单费用
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
func Taobaoalitripieagentrefundnewfillconfirmfee(clt *core.SDKClient, req *flight.TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest, session string) (*flight.TaobaoalitripieagentrefundnewfillconfirmfeeAPIResponse, error) {
	var resp flight.TaobaoalitripieagentrefundnewfillconfirmfeeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
