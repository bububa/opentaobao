package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

/* TaobaoOpentradeQueueQuery
尖货交易排队信息查询
taobao.opentrade.queue.query

尖货交易排队信息查询 */
func TaobaoOpentradeQueueQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeQueueQueryAPIRequest, session string) (*opentrade.TaobaoOpentradeQueueQueryAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeQueueQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
