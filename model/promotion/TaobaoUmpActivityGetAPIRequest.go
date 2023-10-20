package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitygetAPIRequest 查询营销活动 API请求
// taobao.ump.activity.get
//
// 查询营销活动
type TaobaoumpactivitygetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoumpactivitygetRequest 初始化TaobaoumpactivitygetAPIRequest对象
func NewTaobaoumpactivitygetRequest() *TaobaoumpactivitygetAPIRequest {
	return &TaobaoumpactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivitygetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoumpactivitygetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumpactivitygetAPIRequest) GetActId() int64 {
	return r._actId
}
