package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentorderget 【国际机票】查询订单详情
// taobao.alitrip.ie.agent.order.get
//
// 根据订单ID查询订单详情
func Taobaoalitripieagentorderget(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentordergetAPIRequest, session string) (*ieagency.TaobaoalitripieagentordergetAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
