package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详细信息(面单及仓库信息) APIRequest
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息)
*/
type AlibabaOnetouchLogisticsExpressOrderDetailGetRequest struct {
    model.Params

    // 请求参数
    paramQuery   *LogisticsOrderQueryDto 

}

func NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetRequest{
    return &AlibabaOnetouchLogisticsExpressOrderDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.order.detail.get"
}

func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) SetParamQuery(paramQuery *LogisticsOrderQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressOrderDetailGetRequest) GetParamQuery() *LogisticsOrderQueryDto {
    return r.paramQuery
}

