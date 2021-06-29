package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通经销商商品库存设置 API请求
alibaba.lst.ic.stock.items.update

零售通经销商商品库存设置
*/
type AlibabaLstIcStockItemsUpdateRequest struct {
    model.Params
    // 零售通经销商商品库存
    query   *LstItemStockParam
}

// 初始化AlibabaLstIcStockItemsUpdateRequest对象
func NewAlibabaLstIcStockItemsUpdateRequest() *AlibabaLstIcStockItemsUpdateRequest{
    return &AlibabaLstIcStockItemsUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstIcStockItemsUpdateRequest) GetApiMethodName() string {
    return "alibaba.lst.ic.stock.items.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstIcStockItemsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 零售通经销商商品库存
func (r *AlibabaLstIcStockItemsUpdateRequest) SetQuery(query *LstItemStockParam) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaLstIcStockItemsUpdateRequest) GetQuery() *LstItemStockParam {
    return r.query
}
