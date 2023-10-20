package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemListAPIRequest ISV获取口腔标品列表 API请求
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
type AlibabaAlihealthDentalItemListAPIRequest struct {
	model.Params
	// 是否需要测试商品
	_needTestItem bool
}

// NewAlibabaAlihealthDentalItemListRequest 初始化AlibabaAlihealthDentalItemListAPIRequest对象
func NewAlibabaAlihealthDentalItemListRequest() *AlibabaAlihealthDentalItemListAPIRequest {
	return &AlibabaAlihealthDentalItemListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalItemListAPIRequest) Reset() {
	r._needTestItem = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalItemListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNeedTestItem is NeedTestItem Setter
// 是否需要测试商品
func (r *AlibabaAlihealthDentalItemListAPIRequest) SetNeedTestItem(_needTestItem bool) error {
	r._needTestItem = _needTestItem
	r.Set("need_test_item", _needTestItem)
	return nil
}

// GetNeedTestItem NeedTestItem Getter
func (r AlibabaAlihealthDentalItemListAPIRequest) GetNeedTestItem() bool {
	return r._needTestItem
}

var poolAlibabaAlihealthDentalItemListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalItemListRequest()
	},
}

// GetAlibabaAlihealthDentalItemListRequest 从 sync.Pool 获取 AlibabaAlihealthDentalItemListAPIRequest
func GetAlibabaAlihealthDentalItemListAPIRequest() *AlibabaAlihealthDentalItemListAPIRequest {
	return poolAlibabaAlihealthDentalItemListAPIRequest.Get().(*AlibabaAlihealthDentalItemListAPIRequest)
}

// ReleaseAlibabaAlihealthDentalItemListAPIRequest 将 AlibabaAlihealthDentalItemListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalItemListAPIRequest(v *AlibabaAlihealthDentalItemListAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalItemListAPIRequest.Put(v)
}
