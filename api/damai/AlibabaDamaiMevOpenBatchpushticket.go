package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenbatchpushticket 大麦换验平台-第三方对外开放-票单接口batchPushTicket
// alibaba.damai.mev.open.batchpushticket
//
// 批量推送票单
func Alibabadamaimevopenbatchpushticket(clt *core.SDKClient, req *damai.AlibabadamaimevopenbatchpushticketAPIRequest, session string) (*damai.AlibabadamaimevopenbatchpushticketAPIResponse, error) {
	var resp damai.AlibabadamaimevopenbatchpushticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
