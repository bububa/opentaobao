package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpPoItemQueryAPIRequest AliExpress采购单明细查询API API请求
// aliexpress.ascp.po.item.query
//
// AE 供应链仓发 采购单明细查询
type AliexpressAscpPoItemQueryAPIRequest struct {
	model.Params
	// demo
	_purchaseOrderItemQuery *PurchaseOrderItemQueryDto
}

// NewAliexpressAscpPoItemQueryRequest 初始化AliexpressAscpPoItemQueryAPIRequest对象
func NewAliexpressAscpPoItemQueryRequest() *AliexpressAscpPoItemQueryAPIRequest {
	return &AliexpressAscpPoItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpPoItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.po.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpPoItemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPurchaseOrderItemQuery is PurchaseOrderItemQuery Setter
// demo
func (r *AliexpressAscpPoItemQueryAPIRequest) SetPurchaseOrderItemQuery(_purchaseOrderItemQuery *PurchaseOrderItemQueryDto) error {
	r._purchaseOrderItemQuery = _purchaseOrderItemQuery
	r.Set("purchase_order_item_query", _purchaseOrderItemQuery)
	return nil
}

// GetPurchaseOrderItemQuery PurchaseOrderItemQuery Getter
func (r AliexpressAscpPoItemQueryAPIRequest) GetPurchaseOrderItemQuery() *PurchaseOrderItemQueryDto {
	return r._purchaseOrderItemQuery
}
