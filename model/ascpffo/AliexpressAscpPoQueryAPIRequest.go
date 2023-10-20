package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpPoQueryAPIRequest AliExpress采购单查询API API请求
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
type AliexpressAscpPoQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_purchaseOrderQuery *PurchaseOrderQueryDto
}

// NewAliexpressAscpPoQueryRequest 初始化AliexpressAscpPoQueryAPIRequest对象
func NewAliexpressAscpPoQueryRequest() *AliexpressAscpPoQueryAPIRequest {
	return &AliexpressAscpPoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpPoQueryAPIRequest) Reset() {
	r._purchaseOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpPoQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.po.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpPoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpPoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPurchaseOrderQuery is PurchaseOrderQuery Setter
// 系统自动生成
func (r *AliexpressAscpPoQueryAPIRequest) SetPurchaseOrderQuery(_purchaseOrderQuery *PurchaseOrderQueryDto) error {
	r._purchaseOrderQuery = _purchaseOrderQuery
	r.Set("purchase_order_query", _purchaseOrderQuery)
	return nil
}

// GetPurchaseOrderQuery PurchaseOrderQuery Getter
func (r AliexpressAscpPoQueryAPIRequest) GetPurchaseOrderQuery() *PurchaseOrderQueryDto {
	return r._purchaseOrderQuery
}

var poolAliexpressAscpPoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpPoQueryRequest()
	},
}

// GetAliexpressAscpPoQueryRequest 从 sync.Pool 获取 AliexpressAscpPoQueryAPIRequest
func GetAliexpressAscpPoQueryAPIRequest() *AliexpressAscpPoQueryAPIRequest {
	return poolAliexpressAscpPoQueryAPIRequest.Get().(*AliexpressAscpPoQueryAPIRequest)
}

// ReleaseAliexpressAscpPoQueryAPIRequest 将 AliexpressAscpPoQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpPoQueryAPIRequest(v *AliexpressAscpPoQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpPoQueryAPIRequest.Put(v)
}
