package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappUserPhoneGetAPIRequest
获取当前授权用户手机号码 API请求
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码 */
type TaobaoMiniappUserPhoneGetAPIRequest struct {
	model.Params
}

// NewTaobaoMiniappUserPhoneGetRequest 初始化TaobaoMiniappUserPhoneGetAPIRequest对象
func NewTaobaoMiniappUserPhoneGetRequest() *TaobaoMiniappUserPhoneGetAPIRequest {
	return &TaobaoMiniappUserPhoneGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappUserPhoneGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.user.phone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappUserPhoneGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
