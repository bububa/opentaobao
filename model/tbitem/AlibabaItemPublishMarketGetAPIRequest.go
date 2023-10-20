package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitempublishmarketgetAPIRequest 获取商家可发布商品的市场信息 API请求
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
type AlibabaitempublishmarketgetAPIRequest struct {
	model.Params
}

// NewAlibabaitempublishmarketgetRequest 初始化AlibabaitempublishmarketgetAPIRequest对象
func NewAlibabaitempublishmarketgetRequest() *AlibabaitempublishmarketgetAPIRequest {
	return &AlibabaitempublishmarketgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitempublishmarketgetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.market.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitempublishmarketgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitempublishmarketgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
