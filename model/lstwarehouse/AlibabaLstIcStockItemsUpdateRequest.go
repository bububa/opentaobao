package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通经销商商品库存设置 APIRequest
alibaba.lst.ic.stock.items.update

零售通经销商商品库存设置
*/
type AlibabaLstIcStockItemsUpdateRequest struct {
    model.Params

    // 零售通经销商商品库存
    query   *LstItemStockParam 

}

func NewAlibabaLstIcStockItemsUpdateRequest() *AlibabaLstIcStockItemsUpdateRequest{
    return &AlibabaLstIcStockItemsUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstIcStockItemsUpdateRequest) GetApiMethodName() string {
    return "alibaba.lst.ic.stock.items.update"
}

func (r AlibabaLstIcStockItemsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstIcStockItemsUpdateRequest) SetQuery(query *LstItemStockParam) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstIcStockItemsUpdateRequest) GetQuery() *LstItemStockParam {
    return r.query
}

