package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.global.supplier.item.list.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
