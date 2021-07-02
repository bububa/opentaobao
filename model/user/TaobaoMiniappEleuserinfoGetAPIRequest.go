package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappEleuserinfoGetAPIRequest 获取饿了么用户信息 API请求
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
type TaobaoMiniappEleuserinfoGetAPIRequest struct {
	model.Params
	// 怪兽业务方
	_bizScence string
}

// NewTaobaoMiniappEleuserinfoGetRequest 初始化TaobaoMiniappEleuserinfoGetAPIRequest对象
func NewTaobaoMiniappEleuserinfoGetRequest() *TaobaoMiniappEleuserinfoGetAPIRequest {
	return &TaobaoMiniappEleuserinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuserinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizScence Setter
// 怪兽业务方
func (r *TaobaoMiniappEleuserinfoGetAPIRequest) SetBizScence(_bizScence string) error {
	r._bizScence = _bizScence
	r.Set("biz_scence", _bizScence)
	return nil
}

// Get BizScence Getter
func (r TaobaoMiniappEleuserinfoGetAPIRequest) GetBizScence() string {
	return r._bizScence
}
