package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductScoreGetAPIRequest 产品质量分查询 API请求
// alibaba.icbu.product.score.get
//
// 产品质量分查询
type AlibabaIcbuProductScoreGetAPIRequest struct {
	model.Params
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaIcbuProductScoreGetRequest 初始化AlibabaIcbuProductScoreGetAPIRequest对象
func NewAlibabaIcbuProductScoreGetRequest() *AlibabaIcbuProductScoreGetAPIRequest {
	return &AlibabaIcbuProductScoreGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductScoreGetAPIRequest) Reset() {
	r._productId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductScoreGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.score.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductScoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductScoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductScoreGetAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaIcbuProductScoreGetAPIRequest) GetProductId() string {
	return r._productId
}

var poolAlibabaIcbuProductScoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductScoreGetRequest()
	},
}

// GetAlibabaIcbuProductScoreGetRequest 从 sync.Pool 获取 AlibabaIcbuProductScoreGetAPIRequest
func GetAlibabaIcbuProductScoreGetAPIRequest() *AlibabaIcbuProductScoreGetAPIRequest {
	return poolAlibabaIcbuProductScoreGetAPIRequest.Get().(*AlibabaIcbuProductScoreGetAPIRequest)
}

// ReleaseAlibabaIcbuProductScoreGetAPIRequest 将 AlibabaIcbuProductScoreGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductScoreGetAPIRequest(v *AlibabaIcbuProductScoreGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductScoreGetAPIRequest.Put(v)
}
