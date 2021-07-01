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
type AlibabaAscpSalecategoryQueryAPIRequest struct {
    model.Params
    // 货品ID
    _itemId   []int64
}

// 初始化AlibabaAscpSalecategoryQueryAPIRequest对象
func NewAlibabaAscpSalecategoryQueryRequest() *AlibabaAscpSalecategoryQueryAPIRequest{
    return &AlibabaAscpSalecategoryQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.salecategory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 货品ID
func (r *AlibabaAscpSalecategoryQueryAPIRequest) SetItemId(_itemId []int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaAscpSalecategoryQueryAPIRequest) GetItemId() []int64 {
    return r._itemId
}
