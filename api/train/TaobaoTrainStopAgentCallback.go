package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainStopAgentCallback 火车票车次停运信息商家回调
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
func TaobaoTrainStopAgentCallback(clt *core.SDKClient, req *train.TaobaoTrainStopAgentCallbackAPIRequest, session string) (*train.TaobaoTrainStopAgentCallbackAPIResponse, error) {
	var resp train.TaobaoTrainStopAgentCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
