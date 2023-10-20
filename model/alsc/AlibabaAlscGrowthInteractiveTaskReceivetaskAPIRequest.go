package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest 领取任务 API请求
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
type AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest struct {
	model.Params
	// 任务领取入参
	_missionReceiveQuery *MissionReceiveQuery
}

// NewAlibabaalscgrowthinteractivetaskreceivetaskRequest 初始化AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest对象
func NewAlibabaalscgrowthinteractivetaskreceivetaskRequest() *AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest {
	return &AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.receivetask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMissionReceiveQuery is MissionReceiveQuery Setter
// 任务领取入参
func (r *AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest) SetMissionReceiveQuery(_missionReceiveQuery *MissionReceiveQuery) error {
	r._missionReceiveQuery = _missionReceiveQuery
	r.Set("mission_receive_query", _missionReceiveQuery)
	return nil
}

// GetMissionReceiveQuery MissionReceiveQuery Getter
func (r AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest) GetMissionReceiveQuery() *MissionReceiveQuery {
	return r._missionReceiveQuery
}
