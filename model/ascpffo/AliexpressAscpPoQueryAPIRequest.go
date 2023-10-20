package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascppoqueryAPIRequest AliExpress采购单查询API API请求
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
type AliexpressascppoqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_purchaseOrderQuery *PurchaseOrderQueryDto
}

// NewAliexpressascppoqueryRequest 初始化AliexpressascppoqueryAPIRequest对象
func NewAliexpressascppoqueryRequest() *AliexpressascppoqueryAPIRequest {
	return &AliexpressascppoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascppoqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.po.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascppoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascppoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPurchaseOrderQuery is PurchaseOrderQuery Setter
// 系统自动生成
func (r *AliexpressascppoqueryAPIRequest) SetPurchaseOrderQuery(_purchaseOrderQuery *PurchaseOrderQueryDto) error {
	r._purchaseOrderQuery = _purchaseOrderQuery
	r.Set("purchase_order_query", _purchaseOrderQuery)
	return nil
}

// GetPurchaseOrderQuery PurchaseOrderQuery Getter
func (r AliexpressascppoqueryAPIRequest) GetPurchaseOrderQuery() *PurchaseOrderQueryDto {
	return r._purchaseOrderQuery
}
