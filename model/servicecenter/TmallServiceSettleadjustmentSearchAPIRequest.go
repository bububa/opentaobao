package servicecenter

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentSearchAPIRequest) Reset() {
	r._endTime = ""
	r._startTime = ""
	r.Params.ToZero()
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

var poolTmallServiceSettleadjustmentSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentSearchRequest()
	},
}

// GetTmallServiceSettleadjustmentSearchRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentSearchAPIRequest
func GetTmallServiceSettleadjustmentSearchAPIRequest() *TmallServiceSettleadjustmentSearchAPIRequest {
	return poolTmallServiceSettleadjustmentSearchAPIRequest.Get().(*TmallServiceSettleadjustmentSearchAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentSearchAPIRequest 将 TmallServiceSettleadjustmentSearchAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentSearchAPIRequest(v *TmallServiceSettleadjustmentSearchAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentSearchAPIRequest.Put(v)
}
