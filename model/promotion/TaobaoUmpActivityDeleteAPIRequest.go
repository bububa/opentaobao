package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityDeleteAPIRequest 删除营销活动 API请求
// taobao.ump.activity.delete
//
// 删除营销活动。对应的活动详情等将会被全部删除。
type TaobaoUmpActivityDeleteAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpActivityDeleteRequest 初始化TaobaoUmpActivityDeleteAPIRequest对象
func NewTaobaoUmpActivityDeleteRequest() *TaobaoUmpActivityDeleteAPIRequest {
	return &TaobaoUmpActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpActivityDeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpActivityDeleteAPIRequest) GetActId() int64 {
	return r._actId
}
