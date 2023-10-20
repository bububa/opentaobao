package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemvipaddschemagetAPIRequest vip商家发布商品的获取规则接口 API请求
// tmall.item.vip.add.schema.get
//
// 获取vip商家发布商品的规则
type TmallitemvipaddschemagetAPIRequest struct {
	model.Params
}

// NewTmallitemvipaddschemagetRequest 初始化TmallitemvipaddschemagetAPIRequest对象
func NewTmallitemvipaddschemagetRequest() *TmallitemvipaddschemagetAPIRequest {
	return &TmallitemvipaddschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemvipaddschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemvipaddschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemvipaddschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
