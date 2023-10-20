package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAftersaleGetAPIRequest 查询用户售后服务模板 API请求
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
type TaobaoAftersaleGetAPIRequest struct {
	model.Params
}

// NewTaobaoAftersaleGetRequest 初始化TaobaoAftersaleGetAPIRequest对象
func NewTaobaoAftersaleGetRequest() *TaobaoAftersaleGetAPIRequest {
	return &TaobaoAftersaleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAftersaleGetAPIRequest) GetApiMethodName() string {
	return "taobao.aftersale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAftersaleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAftersaleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
