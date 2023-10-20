package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcaseaddproductAPIRequest 批量添加橱窗商品 API请求
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
type AlibabascbpshowcaseaddproductAPIRequest struct {
	model.Params
	// 需要添加的产品ids
	_productIdList []string
}

// NewAlibabascbpshowcaseaddproductRequest 初始化AlibabascbpshowcaseaddproductAPIRequest对象
func NewAlibabascbpshowcaseaddproductRequest() *AlibabascbpshowcaseaddproductAPIRequest {
	return &AlibabascbpshowcaseaddproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcaseaddproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.addproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcaseaddproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcaseaddproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIdList is ProductIdList Setter
// 需要添加的产品ids
func (r *AlibabascbpshowcaseaddproductAPIRequest) SetProductIdList(_productIdList []string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabascbpshowcaseaddproductAPIRequest) GetProductIdList() []string {
	return r._productIdList
}
