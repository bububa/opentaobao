package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappeleuserinfogetAPIRequest 获取饿了么用户信息 API请求
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
type TaobaominiappeleuserinfogetAPIRequest struct {
	model.Params
	// 怪兽业务方
	_bizScence string
}

// NewTaobaominiappeleuserinfogetRequest 初始化TaobaominiappeleuserinfogetAPIRequest对象
func NewTaobaominiappeleuserinfogetRequest() *TaobaominiappeleuserinfogetAPIRequest {
	return &TaobaominiappeleuserinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappeleuserinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuserinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappeleuserinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappeleuserinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizScence is BizScence Setter
// 怪兽业务方
func (r *TaobaominiappeleuserinfogetAPIRequest) SetBizScence(_bizScence string) error {
	r._bizScence = _bizScence
	r.Set("biz_scence", _bizScence)
	return nil
}

// GetBizScence BizScence Getter
func (r TaobaominiappeleuserinfogetAPIRequest) GetBizScence() string {
	return r._bizScence
}
