package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest 本地生活用户增长互动任务平台查询任务 API请求
// alibaba.alsc.growth.interactive.task.querytask
//
// 本地生活用户增长互动任务平台查询任务
type AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest struct {
	model.Params
	// 任务查询入参
	_missionQuery *MissionQuery
}

// NewAlibabaAlscGrowthInteractiveTaskQuerytaskRequest 初始化AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest对象
func NewAlibabaAlscGrowthInteractiveTaskQuerytaskRequest() *AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest {
	return &AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.querytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMissionQuery is MissionQuery Setter
// 任务查询入参
func (r *AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest) SetMissionQuery(_missionQuery *MissionQuery) error {
	r._missionQuery = _missionQuery
	r.Set("mission_query", _missionQuery)
	return nil
}

// GetMissionQuery MissionQuery Getter
func (r AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest) GetMissionQuery() *MissionQuery {
	return r._missionQuery
}
