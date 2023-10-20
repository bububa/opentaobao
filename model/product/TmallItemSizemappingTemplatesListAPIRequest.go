package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsizemappingtemplateslistAPIRequest 获取天猫商品尺码表模板列表 API请求
// tmall.item.sizemapping.templates.list
//
// 获取所有尺码表模板列表。
type TmallitemsizemappingtemplateslistAPIRequest struct {
	model.Params
}

// NewTmallitemsizemappingtemplateslistRequest 初始化TmallitemsizemappingtemplateslistAPIRequest对象
func NewTmallitemsizemappingtemplateslistRequest() *TmallitemsizemappingtemplateslistAPIRequest {
	return &TmallitemsizemappingtemplateslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsizemappingtemplateslistAPIRequest) GetApiMethodName() string {
	return "tmall.item.sizemapping.templates.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsizemappingtemplateslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsizemappingtemplateslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
