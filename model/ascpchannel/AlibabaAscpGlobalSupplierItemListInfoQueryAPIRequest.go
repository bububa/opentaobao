package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest 供应商供品信息查询 API请求
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
type AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest struct {
	model.Params
	// 请求组装
	_supplyProductPageQuery *TopRequest
}

// NewAlibabaascpglobalsupplieritemlistinfoqueryRequest 初始化AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest对象
func NewAlibabaascpglobalsupplieritemlistinfoqueryRequest() *AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest {
	return &AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.global.supplier.item.list.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplyProductPageQuery is SupplyProductPageQuery Setter
// 请求组装
func (r *AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest) SetSupplyProductPageQuery(_supplyProductPageQuery *TopRequest) error {
	r._supplyProductPageQuery = _supplyProductPageQuery
	r.Set("supply_product_page_query", _supplyProductPageQuery)
	return nil
}

// GetSupplyProductPageQuery SupplyProductPageQuery Getter
func (r AlibabaascpglobalsupplieritemlistinfoqueryAPIRequest) GetSupplyProductPageQuery() *TopRequest {
	return r._supplyProductPageQuery
}
