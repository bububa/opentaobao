package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentticketissue 【国际机票】手工出票
// taobao.alitrip.ie.agent.ticket.issue
//
// 代理商手工出票，并回填票号
func Taobaoalitripieagentticketissue(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentticketissueAPIRequest, session string) (*ieagency.TaobaoalitripieagentticketissueAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentticketissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
