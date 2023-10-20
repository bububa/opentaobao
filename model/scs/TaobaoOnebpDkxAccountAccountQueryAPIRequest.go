package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxaccountaccountqueryAPIRequest 万相台账号余额查询 API请求
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
type TaobaoonebpdkxaccountaccountqueryAPIRequest struct {
	model.Params
}

// NewTaobaoonebpdkxaccountaccountqueryRequest 初始化TaobaoonebpdkxaccountaccountqueryAPIRequest对象
func NewTaobaoonebpdkxaccountaccountqueryRequest() *TaobaoonebpdkxaccountaccountqueryAPIRequest {
	return &TaobaoonebpdkxaccountaccountqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxaccountaccountqueryAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.account.account.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxaccountaccountqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxaccountaccountqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
