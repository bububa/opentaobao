package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascppoitemqueryAPIRequest AliExpress采购单明细查询API API请求
// aliexpress.ascp.po.item.query
//
// AE 供应链仓发 采购单明细查询
type AliexpressascppoitemqueryAPIRequest struct {
	model.Params
	// demo
	_purchaseOrderItemQuery *PurchaseOrderItemQueryDto
}

// NewAliexpressascppoitemqueryRequest 初始化AliexpressascppoitemqueryAPIRequest对象
func NewAliexpressascppoitemqueryRequest() *AliexpressascppoitemqueryAPIRequest {
	return &AliexpressascppoitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascppoitemqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.po.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascppoitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascppoitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPurchaseOrderItemQuery is PurchaseOrderItemQuery Setter
// demo
func (r *AliexpressascppoitemqueryAPIRequest) SetPurchaseOrderItemQuery(_purchaseOrderItemQuery *PurchaseOrderItemQueryDto) error {
	r._purchaseOrderItemQuery = _purchaseOrderItemQuery
	r.Set("purchase_order_item_query", _purchaseOrderItemQuery)
	return nil
}

// GetPurchaseOrderItemQuery PurchaseOrderItemQuery Getter
func (r AliexpressascppoitemqueryAPIRequest) GetPurchaseOrderItemQuery() *PurchaseOrderItemQueryDto {
	return r._purchaseOrderItemQuery
}
