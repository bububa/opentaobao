package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainstopagentcallbackAPIRequest 火车票车次停运信息商家回调 API请求
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
type TaobaotrainstopagentcallbackAPIRequest struct {
	model.Params
	// 代理商车次停运信息
	_trainAgentStopInfo *TrainAgentStopInfo
}

// NewTaobaotrainstopagentcallbackRequest 初始化TaobaotrainstopagentcallbackAPIRequest对象
func NewTaobaotrainstopagentcallbackRequest() *TaobaotrainstopagentcallbackAPIRequest {
	return &TaobaotrainstopagentcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainstopagentcallbackAPIRequest) GetApiMethodName() string {
	return "taobao.train.stop.agent.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainstopagentcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainstopagentcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrainAgentStopInfo is TrainAgentStopInfo Setter
// 代理商车次停运信息
func (r *TaobaotrainstopagentcallbackAPIRequest) SetTrainAgentStopInfo(_trainAgentStopInfo *TrainAgentStopInfo) error {
	r._trainAgentStopInfo = _trainAgentStopInfo
	r.Set("train_agent_stop_info", _trainAgentStopInfo)
	return nil
}

// GetTrainAgentStopInfo TrainAgentStopInfo Getter
func (r TaobaotrainstopagentcallbackAPIRequest) GetTrainAgentStopInfo() *TrainAgentStopInfo {
	return r._trainAgentStopInfo
}
