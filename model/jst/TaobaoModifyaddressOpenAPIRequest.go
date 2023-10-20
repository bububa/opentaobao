package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyaddressOpenAPIRequest 淘宝自助修改地址服务开通 API请求
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
type TaobaoModifyaddressOpenAPIRequest struct {
	model.Params
}

// NewTaobaoModifyaddressOpenRequest 初始化TaobaoModifyaddressOpenAPIRequest对象
func NewTaobaoModifyaddressOpenRequest() *TaobaoModifyaddressOpenAPIRequest {
	return &TaobaoModifyaddressOpenAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoModifyaddressOpenAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoModifyaddressOpenAPIRequest) GetApiMethodName() string {
	return "taobao.modifyaddress.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoModifyaddressOpenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoModifyaddressOpenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoModifyaddressOpenAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoModifyaddressOpenRequest()
	},
}

// GetTaobaoModifyaddressOpenRequest 从 sync.Pool 获取 TaobaoModifyaddressOpenAPIRequest
func GetTaobaoModifyaddressOpenAPIRequest() *TaobaoModifyaddressOpenAPIRequest {
	return poolTaobaoModifyaddressOpenAPIRequest.Get().(*TaobaoModifyaddressOpenAPIRequest)
}

// ReleaseTaobaoModifyaddressOpenAPIRequest 将 TaobaoModifyaddressOpenAPIRequest 放入 sync.Pool
func ReleaseTaobaoModifyaddressOpenAPIRequest(v *TaobaoModifyaddressOpenAPIRequest) {
	v.Reset()
	poolTaobaoModifyaddressOpenAPIRequest.Put(v)
}
