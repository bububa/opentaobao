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
type AliexpressAscpItemQueryRequest struct {
    model.Params
    // DTO
    scItemQuery   *ScItemQueryDto
}

// 初始化AliexpressAscpItemQueryRequest对象
func NewAliexpressAscpItemQueryRequest() *AliexpressAscpItemQueryRequest{
    return &AliexpressAscpItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScItemQuery Setter
// DTO
func (r *AliexpressAscpItemQueryRequest) SetScItemQuery(scItemQuery *ScItemQueryDto) error {
    r.scItemQuery = scItemQuery
    r.Set("sc_item_query", scItemQuery)
    return nil
}

// ScItemQuery Getter
func (r AliexpressAscpItemQueryRequest) GetScItemQuery() *ScItemQueryDto {
    return r.scItemQuery
}
