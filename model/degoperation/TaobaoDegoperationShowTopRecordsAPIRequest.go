package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationShowTopRecordsAPIRequest 活动中奖记录 API请求
// taobao.degoperation.show.top.records
//
// 活动中奖记录
type TaobaoDegoperationShowTopRecordsAPIRequest struct {
	model.Params
	// 活动后台配置
	_degAppKey string
	// 活动后台配置
	_degEventKey string
	// 返回数
	_topN int64
}

// NewTaobaoDegoperationShowTopRecordsRequest 初始化TaobaoDegoperationShowTopRecordsAPIRequest对象
func NewTaobaoDegoperationShowTopRecordsRequest() *TaobaoDegoperationShowTopRecordsAPIRequest {
	return &TaobaoDegoperationShowTopRecordsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.show.top.records"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowTopRecordsAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetDegEventKey is DegEventKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowTopRecordsAPIRequest) SetDegEventKey(_degEventKey string) error {
	r._degEventKey = _degEventKey
	r.Set("deg_event_key", _degEventKey)
	return nil
}

// GetDegEventKey DegEventKey Getter
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetDegEventKey() string {
	return r._degEventKey
}

// SetTopN is TopN Setter
// 返回数
func (r *TaobaoDegoperationShowTopRecordsAPIRequest) SetTopN(_topN int64) error {
	r._topN = _topN
	r.Set("top_n", _topN)
	return nil
}

// GetTopN TopN Getter
func (r TaobaoDegoperationShowTopRecordsAPIRequest) GetTopN() int64 {
	return r._topN
}
