package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalitemlistAPIRequest ISV获取口腔标品列表 API请求
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
type AlibabaalihealthdentalitemlistAPIRequest struct {
	model.Params
	// 是否需要测试商品
	_needTestItem bool
}

// NewAlibabaalihealthdentalitemlistRequest 初始化AlibabaalihealthdentalitemlistAPIRequest对象
func NewAlibabaalihealthdentalitemlistRequest() *AlibabaalihealthdentalitemlistAPIRequest {
	return &AlibabaalihealthdentalitemlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalitemlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalitemlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalitemlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNeedTestItem is NeedTestItem Setter
// 是否需要测试商品
func (r *AlibabaalihealthdentalitemlistAPIRequest) SetNeedTestItem(_needTestItem bool) error {
	r._needTestItem = _needTestItem
	r.Set("need_test_item", _needTestItem)
	return nil
}

// GetNeedTestItem NeedTestItem Getter
func (r AlibabaalihealthdentalitemlistAPIRequest) GetNeedTestItem() bool {
	return r._needTestItem
}
