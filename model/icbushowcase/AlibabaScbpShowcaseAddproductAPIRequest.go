package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseAddproductAPIRequest 批量添加橱窗商品 API请求
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
type AlibabaScbpShowcaseAddproductAPIRequest struct {
	model.Params
	// 需要添加的产品ids
	_productIdList []string
}

// NewAlibabaScbpShowcaseAddproductRequest 初始化AlibabaScbpShowcaseAddproductAPIRequest对象
func NewAlibabaScbpShowcaseAddproductRequest() *AlibabaScbpShowcaseAddproductAPIRequest {
	return &AlibabaScbpShowcaseAddproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.addproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProductIdList is ProductIdList Setter
// 需要添加的产品ids
func (r *AlibabaScbpShowcaseAddproductAPIRequest) SetProductIdList(_productIdList []string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetProductIdList() []string {
	return r._productIdList
}
