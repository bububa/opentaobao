package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoModifyaddressOpenAPIRequest
淘宝自助修改地址服务开通 API请求
taobao.modifyaddress.open

商家自助修改地址服务开通 */
type TaobaoModifyaddressOpenAPIRequest struct {
	model.Params
}

// NewTaobaoModifyaddressOpenRequest 初始化TaobaoModifyaddressOpenAPIRequest对象
func NewTaobaoModifyaddressOpenRequest() *TaobaoModifyaddressOpenAPIRequest {
	return &TaobaoModifyaddressOpenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoModifyaddressOpenAPIRequest) GetApiMethodName() string {
	return "taobao.modifyaddress.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoModifyaddressOpenAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
