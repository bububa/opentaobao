package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货品品类查询 APIRequest
alibaba.ascp.salecategory.query

根据货品ID查询对应销售品类ID
*/
type AlibabaAscpSalecategoryQueryRequest struct {
    model.Params

    // 货品ID
    itemId   []int64 

}

func NewAlibabaAscpSalecategoryQueryRequest() *AlibabaAscpSalecategoryQueryRequest{
    return &AlibabaAscpSalecategoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpSalecategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.salecategory.query"
}

func (r AlibabaAscpSalecategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpSalecategoryQueryRequest) SetItemId(itemId []int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaAscpSalecategoryQueryRequest) GetItemId() []int64 {
    return r.itemId
}

