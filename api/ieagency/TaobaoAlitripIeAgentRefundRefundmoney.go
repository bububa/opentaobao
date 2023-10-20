package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundrefundmoney 确认退款
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
func Taobaoalitripieagentrefundrefundmoney(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundrefundmoneyAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundrefundmoneyAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundrefundmoneyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
