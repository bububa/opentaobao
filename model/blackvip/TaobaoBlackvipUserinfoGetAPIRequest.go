package blackvip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoblackvipuserinfogetAPIRequest 88VIP用户信息查询 API请求
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
type TaobaoblackvipuserinfogetAPIRequest struct {
	model.Params
}

// NewTaobaoblackvipuserinfogetRequest 初始化TaobaoblackvipuserinfogetAPIRequest对象
func NewTaobaoblackvipuserinfogetRequest() *TaobaoblackvipuserinfogetAPIRequest {
	return &TaobaoblackvipuserinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoblackvipuserinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.blackvip.userinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoblackvipuserinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoblackvipuserinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
