package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详细信息(面单及仓库信息) API请求
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息)
*/
type AlibabaOnetouchLogisticsExpressOrderDetailGetRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *LogisticsOrderQueryDto
}

// 初始化AlibabaOnetouchLogisticsExpressOrderDetailGetRequest对象
func NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetRequest{
    return &AlibabaOnetouchLogisticsExpressOrderDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.order.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) SetParamQuery(_paramQuery *LogisticsOrderQueryDto) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetParamQuery() *LogisticsOrderQueryDto {
    return r._paramQuery
}
