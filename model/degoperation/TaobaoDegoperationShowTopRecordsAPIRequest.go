package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegoperationshowtoprecordsAPIRequest 活动中奖记录 API请求
// taobao.degoperation.show.top.records
//
// 活动中奖记录
type TaobaodegoperationshowtoprecordsAPIRequest struct {
	model.Params
	// 活动后台配置
	_degAppKey string
	// 活动后台配置
	_degEventKey string
	// 返回数
	_topN int64
}

// NewTaobaodegoperationshowtoprecordsRequest 初始化TaobaodegoperationshowtoprecordsAPIRequest对象
func NewTaobaodegoperationshowtoprecordsRequest() *TaobaodegoperationshowtoprecordsAPIRequest {
	return &TaobaodegoperationshowtoprecordsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.show.top.records"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置
func (r *TaobaodegoperationshowtoprecordsAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetDegEventKey is DegEventKey Setter
// 活动后台配置
func (r *TaobaodegoperationshowtoprecordsAPIRequest) SetDegEventKey(_degEventKey string) error {
	r._degEventKey = _degEventKey
	r.Set("deg_event_key", _degEventKey)
	return nil
}

// GetDegEventKey DegEventKey Getter
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetDegEventKey() string {
	return r._degEventKey
}

// SetTopN is TopN Setter
// 返回数
func (r *TaobaodegoperationshowtoprecordsAPIRequest) SetTopN(_topN int64) error {
	r._topN = _topN
	r.Set("top_n", _topN)
	return nil
}

// GetTopN TopN Getter
func (r TaobaodegoperationshowtoprecordsAPIRequest) GetTopN() int64 {
	return r._topN
}
