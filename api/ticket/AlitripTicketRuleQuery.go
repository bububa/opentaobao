package ticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// Alitripticketrulequery 【门票API2.0】门票规则信息查询接口
// alitrip.ticket.rule.query
//
// 门票规则信息查询接口：返回商家上传的门票规则信息
func Alitripticketrulequery(clt *core.SDKClient, req *ticket.AlitripticketrulequeryAPIRequest, session string) (*ticket.AlitripticketrulequeryAPIResponse, error) {
	var resp ticket.AlitripticketrulequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
