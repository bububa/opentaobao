package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenuidgetAPIRequest 获取授权账号对应的OpenUid API请求
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
type TaobaoopenuidgetAPIRequest struct {
	model.Params
}

// NewTaobaoopenuidgetRequest 初始化TaobaoopenuidgetAPIRequest对象
func NewTaobaoopenuidgetRequest() *TaobaoopenuidgetAPIRequest {
	return &TaobaoopenuidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenuidgetAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenuidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenuidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
