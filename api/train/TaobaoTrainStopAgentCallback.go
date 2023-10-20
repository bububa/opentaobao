package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainstopagentcallback 火车票车次停运信息商家回调
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
func Taobaotrainstopagentcallback(clt *core.SDKClient, req *train.TaobaotrainstopagentcallbackAPIRequest, session string) (*train.TaobaotrainstopagentcallbackAPIResponse, error) {
	var resp train.TaobaotrainstopagentcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
