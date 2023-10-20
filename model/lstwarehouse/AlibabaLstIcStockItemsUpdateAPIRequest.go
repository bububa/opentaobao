package lstwarehouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstIcStockItemsUpdateAPIRequest 零售通经销商商品库存设置 API请求
// alibaba.lst.ic.stock.items.update
//
// 零售通经销商商品库存设置
type AlibabaLstIcStockItemsUpdateAPIRequest struct {
	model.Params
	// 零售通经销商商品库存
	_query *LstItemStockParam
}

// NewAlibabaLstIcStockItemsUpdateRequest 初始化AlibabaLstIcStockItemsUpdateAPIRequest对象
func NewAlibabaLstIcStockItemsUpdateRequest() *AlibabaLstIcStockItemsUpdateAPIRequest {
	return &AlibabaLstIcStockItemsUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstIcStockItemsUpdateAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.ic.stock.items.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 零售通经销商商品库存
func (r *AlibabaLstIcStockItemsUpdateAPIRequest) SetQuery(_query *LstItemStockParam) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetQuery() *LstItemStockParam {
	return r._query
}

var poolAlibabaLstIcStockItemsUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstIcStockItemsUpdateRequest()
	},
}

// GetAlibabaLstIcStockItemsUpdateRequest 从 sync.Pool 获取 AlibabaLstIcStockItemsUpdateAPIRequest
func GetAlibabaLstIcStockItemsUpdateAPIRequest() *AlibabaLstIcStockItemsUpdateAPIRequest {
	return poolAlibabaLstIcStockItemsUpdateAPIRequest.Get().(*AlibabaLstIcStockItemsUpdateAPIRequest)
}

// ReleaseAlibabaLstIcStockItemsUpdateAPIRequest 将 AlibabaLstIcStockItemsUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstIcStockItemsUpdateAPIRequest(v *AlibabaLstIcStockItemsUpdateAPIRequest) {
	v.Reset()
	poolAlibabaLstIcStockItemsUpdateAPIRequest.Put(v)
}
