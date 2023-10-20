package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoaftersalegetAPIRequest 查询用户售后服务模板 API请求
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
type TaobaoaftersalegetAPIRequest struct {
	model.Params
}

// NewTaobaoaftersalegetRequest 初始化TaobaoaftersalegetAPIRequest对象
func NewTaobaoaftersalegetRequest() *TaobaoaftersalegetAPIRequest {
	return &TaobaoaftersalegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoaftersalegetAPIRequest) GetApiMethodName() string {
	return "taobao.aftersale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoaftersalegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoaftersalegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
