package ascpffo

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpPoItemQueryAPIRequest) Reset() {
	r._purchaseOrderItemQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpPoItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.po.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpPoItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpPoItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAliexpressAscpPoItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpPoItemQueryRequest()
	},
}

// GetAliexpressAscpPoItemQueryRequest 从 sync.Pool 获取 AliexpressAscpPoItemQueryAPIRequest
func GetAliexpressAscpPoItemQueryAPIRequest() *AliexpressAscpPoItemQueryAPIRequest {
	return poolAliexpressAscpPoItemQueryAPIRequest.Get().(*AliexpressAscpPoItemQueryAPIRequest)
}

// ReleaseAliexpressAscpPoItemQueryAPIRequest 将 AliexpressAscpPoItemQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpPoItemQueryAPIRequest(v *AliexpressAscpPoItemQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpPoItemQueryAPIRequest.Put(v)
}
