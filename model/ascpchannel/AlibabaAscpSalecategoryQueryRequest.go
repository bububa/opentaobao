package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货品品类查询 API请求
alibaba.ascp.salecategory.query

根据货品ID查询对应销售品类ID
*/
type AlibabaAscpSalecategoryQueryRequest struct {
    model.Params
    // 货品ID
    itemId   []int64
}

// 初始化AlibabaAscpSalecategoryQueryRequest对象
func NewAlibabaAscpSalecategoryQueryRequest() *AlibabaAscpSalecategoryQueryRequest{
    return &AlibabaAscpSalecategoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpSalecategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.salecategory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpSalecategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 货品ID
func (r *AlibabaAscpSalecategoryQueryRequest) SetItemId(itemId []int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaAscpSalecategoryQueryRequest) GetItemId() []int64 {
    return r.itemId
}
