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
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *LogisticsOrderQueryDto
}

// 初始化AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest{
    return &AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.order.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) SetParamQuery(_paramQuery *LogisticsOrderQueryDto) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetParamQuery() *LogisticsOrderQueryDto {
    return r._paramQuery
}
