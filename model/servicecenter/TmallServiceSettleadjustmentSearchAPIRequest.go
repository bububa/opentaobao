package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentSearchAPIRequest 服务商15分钟获取数据 API请求
// tmall.service.settleadjustment.search
//
// 天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
type TmallServiceSettleadjustmentSearchAPIRequest struct {
	model.Params
	// 结束时间
	_endTime string
	// 开始时间
	_startTime string
}

// NewTmallServiceSettleadjustmentSearchRequest 初始化TmallServiceSettleadjustmentSearchAPIRequest对象
func NewTmallServiceSettleadjustmentSearchRequest() *TmallServiceSettleadjustmentSearchAPIRequest {
	return &TmallServiceSettleadjustmentSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentSearchAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TmallServiceSettleadjustmentSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TmallServiceSettleadjustmentSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TmallServiceSettleadjustmentSearchAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TmallServiceSettleadjustmentSearchAPIRequest) GetStartTime() string {
	return r._startTime
}
