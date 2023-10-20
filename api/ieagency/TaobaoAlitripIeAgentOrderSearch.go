package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentordersearch 【国际机票】订单列表查询
// taobao.alitrip.ie.agent.order.search
//
// 根据指定条件查询国际机票订单列表
func Taobaoalitripieagentordersearch(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentordersearchAPIRequest, session string) (*ieagency.TaobaoalitripieagentordersearchAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
