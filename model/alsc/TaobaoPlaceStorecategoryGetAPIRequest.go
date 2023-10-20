package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorecategorygetAPIRequest 获取门店类目信息 API请求
// taobao.place.storecategory.get
//
// 获取门店类目信息
type TaobaoplacestorecategorygetAPIRequest struct {
	model.Params
}

// NewTaobaoplacestorecategorygetRequest 初始化TaobaoplacestorecategorygetAPIRequest对象
func NewTaobaoplacestorecategorygetRequest() *TaobaoplacestorecategorygetAPIRequest {
	return &TaobaoplacestorecategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorecategorygetAPIRequest) GetApiMethodName() string {
	return "taobao.place.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorecategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorecategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
