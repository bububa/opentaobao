package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-市 APIRequest
alibaba.onetouch.logistics.express.address.city.list

四级地址库-市
*/
type AlibabaOnetouchLogisticsExpressAddressCityListRequest struct {
    model.Params

    // 请求参数
    paramQuery   *AddressQueryDto 

}

func NewAlibabaOnetouchLogisticsExpressAddressCityListRequest() *AlibabaOnetouchLogisticsExpressAddressCityListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressCityListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.city.list"
}

func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressAddressCityListRequest) SetParamQuery(paramQuery *AddressQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetParamQuery() *AddressQueryDto {
    return r.paramQuery
}

