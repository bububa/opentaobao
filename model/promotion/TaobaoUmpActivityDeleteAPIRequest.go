package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitydeleteAPIRequest 删除营销活动 API请求
// taobao.ump.activity.delete
//
// 删除营销活动。对应的活动详情等将会被全部删除。
type TaobaoumpactivitydeleteAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoumpactivitydeleteRequest 初始化TaobaoumpactivitydeleteAPIRequest对象
func NewTaobaoumpactivitydeleteRequest() *TaobaoumpactivitydeleteAPIRequest {
	return &TaobaoumpactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivitydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoumpactivitydeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumpactivitydeleteAPIRequest) GetActId() int64 {
	return r._actId
}
