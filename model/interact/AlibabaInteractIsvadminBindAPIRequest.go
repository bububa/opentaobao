package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvadminBindAPIRequest 创建及绑定互动实例 API请求
// alibaba.interact.isvadmin.bind
//
// 创建互动实例，并绑定奖池
type AlibabaInteractIsvadminBindAPIRequest struct {
	model.Params
	// 描述信息
	_interactDescription string
	// 互动实例名称
	_instanceName string
	// 互动开始时间
	_startTime string
	// 互动结束时间
	_endTime string
	// 奖池ID
	_lotteryCode string
}

// NewAlibabaInteractIsvadminBindRequest 初始化AlibabaInteractIsvadminBindAPIRequest对象
func NewAlibabaInteractIsvadminBindRequest() *AlibabaInteractIsvadminBindAPIRequest {
	return &AlibabaInteractIsvadminBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminBindAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isvadmin.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInteractDescription is InteractDescription Setter
// 描述信息
func (r *AlibabaInteractIsvadminBindAPIRequest) SetInteractDescription(_interactDescription string) error {
	r._interactDescription = _interactDescription
	r.Set("interact_description", _interactDescription)
	return nil
}

// GetInteractDescription InteractDescription Getter
func (r AlibabaInteractIsvadminBindAPIRequest) GetInteractDescription() string {
	return r._interactDescription
}

// SetInstanceName is InstanceName Setter
// 互动实例名称
func (r *AlibabaInteractIsvadminBindAPIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("instance_name", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r AlibabaInteractIsvadminBindAPIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetStartTime is StartTime Setter
// 互动开始时间
func (r *AlibabaInteractIsvadminBindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaInteractIsvadminBindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 互动结束时间
func (r *AlibabaInteractIsvadminBindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaInteractIsvadminBindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetLotteryCode is LotteryCode Setter
// 奖池ID
func (r *AlibabaInteractIsvadminBindAPIRequest) SetLotteryCode(_lotteryCode string) error {
	r._lotteryCode = _lotteryCode
	r.Set("lottery_code", _lotteryCode)
	return nil
}

// GetLotteryCode LotteryCode Getter
func (r AlibabaInteractIsvadminBindAPIRequest) GetLotteryCode() string {
	return r._lotteryCode
}
