package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemaddsimpleschemagetAPIRequest 天猫发布商品规则获取 API请求
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
type TmallitemaddsimpleschemagetAPIRequest struct {
	model.Params
}

// NewTmallitemaddsimpleschemagetRequest 初始化TmallitemaddsimpleschemagetAPIRequest对象
func NewTmallitemaddsimpleschemagetRequest() *TmallitemaddsimpleschemagetAPIRequest {
	return &TmallitemaddsimpleschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemaddsimpleschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.add.simpleschema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemaddsimpleschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemaddsimpleschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
