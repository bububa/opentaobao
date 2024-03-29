package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest 领取任务 API请求
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
type AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest struct {
	model.Params
	// 任务领取入参
	_missionReceiveQuery *MissionReceiveQuery
}

// NewAlibabaAlscGrowthInteractiveTaskReceivetaskRequest 初始化AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest对象
func NewAlibabaAlscGrowthInteractiveTaskReceivetaskRequest() *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest {
	return &AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) Reset() {
	r._missionReceiveQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.receivetask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMissionReceiveQuery is MissionReceiveQuery Setter
// 任务领取入参
func (r *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) SetMissionReceiveQuery(_missionReceiveQuery *MissionReceiveQuery) error {
	r._missionReceiveQuery = _missionReceiveQuery
	r.Set("mission_receive_query", _missionReceiveQuery)
	return nil
}

// GetMissionReceiveQuery MissionReceiveQuery Getter
func (r AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) GetMissionReceiveQuery() *MissionReceiveQuery {
	return r._missionReceiveQuery
}

var poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscGrowthInteractiveTaskReceivetaskRequest()
	},
}

// GetAlibabaAlscGrowthInteractiveTaskReceivetaskRequest 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest
func GetAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest() *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest {
	return poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest.Get().(*AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest 将 AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest(v *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest.Put(v)
}
