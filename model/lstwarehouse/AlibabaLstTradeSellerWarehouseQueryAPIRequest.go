package lstwarehouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerWarehouseQueryAPIRequest 供应商-本云商家-仓库查询接口 API请求
// alibaba.lst.trade.seller.warehouse.query
//
// 查询本地云仓商家的仓库
type AlibabaLstTradeSellerWarehouseQueryAPIRequest struct {
	model.Params
	// 入参
	_warehouseQueryParam *WarehouseQueryParam
}

// NewAlibabaLstTradeSellerWarehouseQueryRequest 初始化AlibabaLstTradeSellerWarehouseQueryAPIRequest对象
func NewAlibabaLstTradeSellerWarehouseQueryRequest() *AlibabaLstTradeSellerWarehouseQueryAPIRequest {
	return &AlibabaLstTradeSellerWarehouseQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeSellerWarehouseQueryAPIRequest) Reset() {
	r._warehouseQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerWarehouseQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.warehouse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerWarehouseQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerWarehouseQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseQueryParam is WarehouseQueryParam Setter
// 入参
func (r *AlibabaLstTradeSellerWarehouseQueryAPIRequest) SetWarehouseQueryParam(_warehouseQueryParam *WarehouseQueryParam) error {
	r._warehouseQueryParam = _warehouseQueryParam
	r.Set("warehouse_query_param", _warehouseQueryParam)
	return nil
}

// GetWarehouseQueryParam WarehouseQueryParam Getter
func (r AlibabaLstTradeSellerWarehouseQueryAPIRequest) GetWarehouseQueryParam() *WarehouseQueryParam {
	return r._warehouseQueryParam
}

var poolAlibabaLstTradeSellerWarehouseQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeSellerWarehouseQueryRequest()
	},
}

// GetAlibabaLstTradeSellerWarehouseQueryRequest 从 sync.Pool 获取 AlibabaLstTradeSellerWarehouseQueryAPIRequest
func GetAlibabaLstTradeSellerWarehouseQueryAPIRequest() *AlibabaLstTradeSellerWarehouseQueryAPIRequest {
	return poolAlibabaLstTradeSellerWarehouseQueryAPIRequest.Get().(*AlibabaLstTradeSellerWarehouseQueryAPIRequest)
}

// ReleaseAlibabaLstTradeSellerWarehouseQueryAPIRequest 将 AlibabaLstTradeSellerWarehouseQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeSellerWarehouseQueryAPIRequest(v *AlibabaLstTradeSellerWarehouseQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeSellerWarehouseQueryAPIRequest.Put(v)
}
