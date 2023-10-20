package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest 任务领奖 API请求
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
type AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest struct {
	model.Params
	// 任务领奖入参
	_rewardReceiveQuery *RewardReceiveQuery
}

// NewAlibabaAlscGrowthInteractiveTaskReceivetaskprizeRequest 初始化AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest对象
func NewAlibabaAlscGrowthInteractiveTaskReceivetaskprizeRequest() *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest {
	return &AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) Reset() {
	r._rewardReceiveQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.receivetaskprize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRewardReceiveQuery is RewardReceiveQuery Setter
// 任务领奖入参
func (r *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) SetRewardReceiveQuery(_rewardReceiveQuery *RewardReceiveQuery) error {
	r._rewardReceiveQuery = _rewardReceiveQuery
	r.Set("reward_receive_query", _rewardReceiveQuery)
	return nil
}

// GetRewardReceiveQuery RewardReceiveQuery Getter
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) GetRewardReceiveQuery() *RewardReceiveQuery {
	return r._rewardReceiveQuery
}

var poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscGrowthInteractiveTaskReceivetaskprizeRequest()
	},
}

// GetAlibabaAlscGrowthInteractiveTaskReceivetaskprizeRequest 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest
func GetAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest() *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest {
	return poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest.Get().(*AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest 将 AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest(v *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest.Put(v)
}
