package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradequeuequery 尖货交易排队信息查询
// taobao.opentrade.queue.query
//
// 尖货交易排队信息查询
func Taobaoopentradequeuequery(clt *core.SDKClient, req *opentrade.TaobaoopentradequeuequeryAPIRequest, session string) (*opentrade.TaobaoopentradequeuequeryAPIResponse, error) {
	var resp opentrade.TaobaoopentradequeuequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
