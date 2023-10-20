package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundnewgetlist 新查询退票申请单列表
// taobao.alitrip.ie.agent.refund.new.getlist
//
// 查询商家国际机票退票申请单列表
func Taobaoalitripieagentrefundnewgetlist(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundnewgetlistAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundnewgetlistAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundnewgetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
