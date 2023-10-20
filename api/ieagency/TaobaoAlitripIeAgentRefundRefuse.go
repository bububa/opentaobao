package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundrefuse 拒绝退票申请
// taobao.alitrip.ie.agent.refund.refuse
//
// 卖家拒绝退票退票申请
func Taobaoalitripieagentrefundrefuse(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundrefuseAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundrefuseAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
