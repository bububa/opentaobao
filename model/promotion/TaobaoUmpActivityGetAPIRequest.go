package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityGetAPIRequest 查询营销活动 API请求
// taobao.ump.activity.get
//
// 查询营销活动
type TaobaoUmpActivityGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpActivityGetRequest 初始化TaobaoUmpActivityGetAPIRequest对象
func NewTaobaoUmpActivityGetRequest() *TaobaoUmpActivityGetAPIRequest {
	return &TaobaoUmpActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpActivityGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpActivityGetAPIRequest) GetActId() int64 {
	return r._actId
}
