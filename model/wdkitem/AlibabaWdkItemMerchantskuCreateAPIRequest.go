package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuCreateAPIRequest 商家商品信息新建 API请求
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
type AlibabaWdkItemMerchantskuCreateAPIRequest struct {
	model.Params
	// 新建商品参数，是个json字符串
	_params string
}

// NewAlibabaWdkItemMerchantskuCreateRequest 初始化AlibabaWdkItemMerchantskuCreateAPIRequest对象
func NewAlibabaWdkItemMerchantskuCreateRequest() *AlibabaWdkItemMerchantskuCreateAPIRequest {
	return &AlibabaWdkItemMerchantskuCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMerchantskuCreateAPIRequest) Reset() {
	r._params = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantsku.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// 新建商品参数，是个json字符串
func (r *AlibabaWdkItemMerchantskuCreateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaWdkItemMerchantskuCreateAPIRequest) GetParams() string {
	return r._params
}

var poolAlibabaWdkItemMerchantskuCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMerchantskuCreateRequest()
	},
}

// GetAlibabaWdkItemMerchantskuCreateRequest 从 sync.Pool 获取 AlibabaWdkItemMerchantskuCreateAPIRequest
func GetAlibabaWdkItemMerchantskuCreateAPIRequest() *AlibabaWdkItemMerchantskuCreateAPIRequest {
	return poolAlibabaWdkItemMerchantskuCreateAPIRequest.Get().(*AlibabaWdkItemMerchantskuCreateAPIRequest)
}

// ReleaseAlibabaWdkItemMerchantskuCreateAPIRequest 将 AlibabaWdkItemMerchantskuCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMerchantskuCreateAPIRequest(v *AlibabaWdkItemMerchantskuCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMerchantskuCreateAPIRequest.Put(v)
}
