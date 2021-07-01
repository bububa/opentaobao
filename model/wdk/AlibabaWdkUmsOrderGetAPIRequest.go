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
type AlibabaWdkUmsOrderGetAPIRequest struct {
    model.Params
    // 查询单据的dto
    _queryErpbillDto   *QueryErpBillDto
}

// 初始化AlibabaWdkUmsOrderGetAPIRequest对象
func NewAlibabaWdkUmsOrderGetRequest() *AlibabaWdkUmsOrderGetAPIRequest{
    return &AlibabaWdkUmsOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryErpbillDto Setter
// 查询单据的dto
func (r *AlibabaWdkUmsOrderGetAPIRequest) SetQueryErpbillDto(_queryErpbillDto *QueryErpBillDto) error {
    r._queryErpbillDto = _queryErpbillDto
    r.Set("query_erpbill_dto", _queryErpbillDto)
    return nil
}

// QueryErpbillDto Getter
func (r AlibabaWdkUmsOrderGetAPIRequest) GetQueryErpbillDto() *QueryErpBillDto {
    return r._queryErpbillDto
}
