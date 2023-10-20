package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaoalitripieagentshoppingpush 国际机票大卖家Shopping推送
// taobao.alitrip.ie.agent.shopping.push
//
// 用于国际机票大卖家主动推送Shopping结果更新缓存报价。
func Taobaoalitripieagentshoppingpush(clt *core.SDKClient, req *flight.TaobaoalitripieagentshoppingpushAPIRequest, session string) (*flight.TaobaoalitripieagentshoppingpushAPIResponse, error) {
	var resp flight.TaobaoalitripieagentshoppingpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
