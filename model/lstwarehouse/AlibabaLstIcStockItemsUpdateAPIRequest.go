package lstwarehouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.ic.stock.items.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 零售通经销商商品库存
func (r *AlibabaLstIcStockItemsUpdateAPIRequest) SetQuery(_query *LstItemStockParam) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaLstIcStockItemsUpdateAPIRequest) GetQuery() *LstItemStockParam {
	return r._query
}
