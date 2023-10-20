package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest 供应商供品信息查询 API请求
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
type AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest struct {
	model.Params
	// 请求组装
	_supplyProductPageQuery *TopRequest
}

// NewAlibabaAscpGlobalSupplierItemListInfoQueryRequest 初始化AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest对象
func NewAlibabaAscpGlobalSupplierItemListInfoQueryRequest() *AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest {
	return &AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) Reset() {
	r._supplyProductPageQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.global.supplier.item.list.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplyProductPageQuery is SupplyProductPageQuery Setter
// 请求组装
func (r *AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) SetSupplyProductPageQuery(_supplyProductPageQuery *TopRequest) error {
	r._supplyProductPageQuery = _supplyProductPageQuery
	r.Set("supply_product_page_query", _supplyProductPageQuery)
	return nil
}

// GetSupplyProductPageQuery SupplyProductPageQuery Getter
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetSupplyProductPageQuery() *TopRequest {
	return r._supplyProductPageQuery
}

var poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpGlobalSupplierItemListInfoQueryRequest()
	},
}

// GetAlibabaAscpGlobalSupplierItemListInfoQueryRequest 从 sync.Pool 获取 AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest
func GetAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest() *AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest {
	return poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest.Get().(*AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest)
}

// ReleaseAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest 将 AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest(v *AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest.Put(v)
}
