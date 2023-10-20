package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumppromotionglobaldiscountgetAPIRequest 获取卖家最低折扣 API请求
// taobao.ump.promotion.global.discount.get
//
// 提供卖家最低折扣查询功能
type TaobaoumppromotionglobaldiscountgetAPIRequest struct {
	model.Params
}

// NewTaobaoumppromotionglobaldiscountgetRequest 初始化TaobaoumppromotionglobaldiscountgetAPIRequest对象
func NewTaobaoumppromotionglobaldiscountgetRequest() *TaobaoumppromotionglobaldiscountgetAPIRequest {
	return &TaobaoumppromotionglobaldiscountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumppromotionglobaldiscountgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.global.discount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumppromotionglobaldiscountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumppromotionglobaldiscountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
