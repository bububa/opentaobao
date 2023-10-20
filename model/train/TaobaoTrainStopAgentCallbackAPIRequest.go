package train

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainStopAgentCallbackAPIRequest) Reset() {
	r._trainAgentStopInfo = nil
	r.Params.ToZero()
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

var poolTaobaoTrainStopAgentCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainStopAgentCallbackRequest()
	},
}

// GetTaobaoTrainStopAgentCallbackRequest 从 sync.Pool 获取 TaobaoTrainStopAgentCallbackAPIRequest
func GetTaobaoTrainStopAgentCallbackAPIRequest() *TaobaoTrainStopAgentCallbackAPIRequest {
	return poolTaobaoTrainStopAgentCallbackAPIRequest.Get().(*TaobaoTrainStopAgentCallbackAPIRequest)
}

// ReleaseTaobaoTrainStopAgentCallbackAPIRequest 将 TaobaoTrainStopAgentCallbackAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainStopAgentCallbackAPIRequest(v *TaobaoTrainStopAgentCallbackAPIRequest) {
	v.Reset()
	poolTaobaoTrainStopAgentCallbackAPIRequest.Put(v)
}
