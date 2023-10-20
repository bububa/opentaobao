package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundnewgetdetail 查询申请单详情(新版)
// taobao.alitrip.ie.agent.refund.new.getdetail
//
// 查询申请单详情
func Taobaoalitripieagentrefundnewgetdetail(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundnewgetdetailAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundnewgetdetailAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundnewgetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
