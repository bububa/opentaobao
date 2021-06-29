package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-区域 APIRequest
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区
*/
type AlibabaOnetouchLogisticsExpressAddressDivisionListRequest struct {
    model.Params

    // 请求参数
    paramQuery   *AddressQueryDto 

}

func NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressDivisionListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.division.list"
}

func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) SetParamQuery(paramQuery *AddressQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetParamQuery() *AddressQueryDto {
    return r.paramQuery
}

