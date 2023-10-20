package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundget 获取退票申请详情
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
func Taobaoalitripieagentrefundget(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundgetAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundgetAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
