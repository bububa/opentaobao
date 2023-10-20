package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainStopAgentCallbackAPIRequest 火车票车次停运信息商家回调 API请求
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
type TaobaoTrainStopAgentCallbackAPIRequest struct {
	model.Params
	// 代理商车次停运信息
	_trainAgentStopInfo *TrainAgentStopInfo
}

// NewTaobaoTrainStopAgentCallbackRequest 初始化TaobaoTrainStopAgentCallbackAPIRequest对象
func NewTaobaoTrainStopAgentCallbackRequest() *TaobaoTrainStopAgentCallbackAPIRequest {
	return &TaobaoTrainStopAgentCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainStopAgentCallbackAPIRequest) GetApiMethodName() string {
	return "taobao.train.stop.agent.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainStopAgentCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainStopAgentCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrainAgentStopInfo is TrainAgentStopInfo Setter
// 代理商车次停运信息
func (r *TaobaoTrainStopAgentCallbackAPIRequest) SetTrainAgentStopInfo(_trainAgentStopInfo *TrainAgentStopInfo) error {
	r._trainAgentStopInfo = _trainAgentStopInfo
	r.Set("train_agent_stop_info", _trainAgentStopInfo)
	return nil
}

// GetTrainAgentStopInfo TrainAgentStopInfo Getter
func (r TaobaoTrainStopAgentCallbackAPIRequest) GetTrainAgentStopInfo() *TrainAgentStopInfo {
	return r._trainAgentStopInfo
}
