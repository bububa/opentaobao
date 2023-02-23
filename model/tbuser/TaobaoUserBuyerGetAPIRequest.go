package tbuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserBuyerGetAPIRequest 查询买家信息API API请求
// taobao.user.buyer.get
//
// 查询买家信息API，只能买家类应用调用。
type TaobaoUserBuyerGetAPIRequest struct {
	model.Params
	// 只返回nick, avatar参数
	_fields string
}

// NewTaobaoUserBuyerGetRequest 初始化TaobaoUserBuyerGetAPIRequest对象
func NewTaobaoUserBuyerGetRequest() *TaobaoUserBuyerGetAPIRequest {
	return &TaobaoUserBuyerGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserBuyerGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.buyer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserBuyerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserBuyerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 只返回nick, avatar参数
func (r *TaobaoUserBuyerGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoUserBuyerGetAPIRequest) GetFields() string {
	return r._fields
}
