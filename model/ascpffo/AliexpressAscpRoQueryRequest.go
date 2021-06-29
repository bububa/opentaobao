package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单查询API API请求
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口
*/
type AliexpressAscpRoQueryRequest struct {
    model.Params
    // dto
    returnOrderQuery   *ReturnOrderQueryDto
}

// 初始化AliexpressAscpRoQueryRequest对象
func NewAliexpressAscpRoQueryRequest() *AliexpressAscpRoQueryRequest{
    return &AliexpressAscpRoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpRoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ro.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpRoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReturnOrderQuery Setter
// dto
func (r *AliexpressAscpRoQueryRequest) SetReturnOrderQuery(returnOrderQuery *ReturnOrderQueryDto) error {
    r.returnOrderQuery = returnOrderQuery
    r.Set("return_order_query", returnOrderQuery)
    return nil
}

// ReturnOrderQuery Getter
func (r AliexpressAscpRoQueryRequest) GetReturnOrderQuery() *ReturnOrderQueryDto {
    return r.returnOrderQuery
}
