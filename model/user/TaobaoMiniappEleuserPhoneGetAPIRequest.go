package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappEleuserPhoneGetAPIRequest
获取饿了么用户信息 API请求
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息 */
type TaobaoMiniappEleuserPhoneGetAPIRequest struct {
	model.Params
}

// NewTaobaoMiniappEleuserPhoneGetRequest 初始化TaobaoMiniappEleuserPhoneGetAPIRequest对象
func NewTaobaoMiniappEleuserPhoneGetRequest() *TaobaoMiniappEleuserPhoneGetAPIRequest {
	return &TaobaoMiniappEleuserPhoneGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserPhoneGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.eleuser.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserPhoneGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
