package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress货品查询查询API API请求
aliexpress.ascp.item.query

AE货品查询API
*/
type AliexpressAscpItemQueryAPIRequest struct {
    model.Params
    // DTO
    _scItemQuery   *ScItemQueryDTO
}

// 初始化AliexpressAscpItemQueryAPIRequest对象
func NewAliexpressAscpItemQueryRequest() *AliexpressAscpItemQueryAPIRequest{
    return &AliexpressAscpItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpItemQueryAPIRequest) GetApiMethodName() string {
    return "aliexpress.ascp.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScItemQuery Setter
// DTO
func (r *AliexpressAscpItemQueryAPIRequest) SetScItemQuery(_scItemQuery *ScItemQueryDTO) error {
    r._scItemQuery = _scItemQuery
    r.Set("sc_item_query", _scItemQuery)
    return nil
}

// ScItemQuery Getter
func (r AliexpressAscpItemQueryAPIRequest) GetScItemQuery() *ScItemQueryDTO {
    return r._scItemQuery
}
