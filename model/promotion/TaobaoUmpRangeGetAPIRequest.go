package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumprangegetAPIRequest 查询活动范围 API请求
// taobao.ump.range.get
//
// 查询某个卖家所有参加或者不参加某项活动的物品
type TaobaoumprangegetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoumprangegetRequest 初始化TaobaoumprangegetAPIRequest对象
func NewTaobaoumprangegetRequest() *TaobaoumprangegetAPIRequest {
	return &TaobaoumprangegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumprangegetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumprangegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumprangegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoumprangegetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumprangegetAPIRequest) GetActId() int64 {
	return r._actId
}
