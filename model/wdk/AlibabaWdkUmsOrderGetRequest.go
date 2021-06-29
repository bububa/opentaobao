package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询店仓作业单据清单 （库存对账辅助）-回流单 API请求
alibaba.wdk.ums.order.get

查询店仓作业单据清单 （库存对账辅助）-回流单
*/
type AlibabaWdkUmsOrderGetRequest struct {
    model.Params
    // 查询单据的dto
    queryErpbillDto   *QueryErpBillDto
}

// 初始化AlibabaWdkUmsOrderGetRequest对象
func NewAlibabaWdkUmsOrderGetRequest() *AlibabaWdkUmsOrderGetRequest{
    return &AlibabaWdkUmsOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryErpbillDto Setter
// 查询单据的dto
func (r *AlibabaWdkUmsOrderGetRequest) SetQueryErpbillDto(queryErpbillDto *QueryErpBillDto) error {
    r.queryErpbillDto = queryErpbillDto
    r.Set("query_erpbill_dto", queryErpbillDto)
    return nil
}

// QueryErpbillDto Getter
func (r AlibabaWdkUmsOrderGetRequest) GetQueryErpbillDto() *QueryErpBillDto {
    return r.queryErpbillDto
}
