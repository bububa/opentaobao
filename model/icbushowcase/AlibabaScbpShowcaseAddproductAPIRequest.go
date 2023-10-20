package icbushowcase

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseAddproductAPIRequest) Reset() {
	r._productIdList = r._productIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.addproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseAddproductAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpShowcaseAddproductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseAddproductRequest()
	},
}

// GetAlibabaScbpShowcaseAddproductRequest 从 sync.Pool 获取 AlibabaScbpShowcaseAddproductAPIRequest
func GetAlibabaScbpShowcaseAddproductAPIRequest() *AlibabaScbpShowcaseAddproductAPIRequest {
	return poolAlibabaScbpShowcaseAddproductAPIRequest.Get().(*AlibabaScbpShowcaseAddproductAPIRequest)
}

// ReleaseAlibabaScbpShowcaseAddproductAPIRequest 将 AlibabaScbpShowcaseAddproductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseAddproductAPIRequest(v *AlibabaScbpShowcaseAddproductAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseAddproductAPIRequest.Put(v)
}
