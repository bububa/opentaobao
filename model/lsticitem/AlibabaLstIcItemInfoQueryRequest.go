package lsticitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 API请求
alibaba.lst.ic.item.info.query

查询商品信息
*/
type AlibabaLstIcItemInfoQueryRequest struct {
    model.Params
    // 零售通商品查询参数
    _query   *LstItemListParam
}

// 初始化AlibabaLstIcItemInfoQueryRequest对象
func NewAlibabaLstIcItemInfoQueryRequest() *AlibabaLstIcItemInfoQueryRequest{
    return &AlibabaLstIcItemInfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstIcItemInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.ic.item.info.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstIcItemInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 零售通商品查询参数
func (r *AlibabaLstIcItemInfoQueryRequest) SetQuery(_query *LstItemListParam) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstIcItemInfoQueryRequest) GetQuery() *LstItemListParam {
    return r._query
}
