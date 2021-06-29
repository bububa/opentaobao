package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-街道 APIRequest
alibaba.onetouch.logistics.express.address.street.list

四级地址库-街道模糊查询
*/
type AlibabaOnetouchLogisticsExpressAddressStreetListRequest struct {
    model.Params

    // 请求参数
    paramQuery   *AddressQueryDto 

}

func NewAlibabaOnetouchLogisticsExpressAddressStreetListRequest() *AlibabaOnetouchLogisticsExpressAddressStreetListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressStreetListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.street.list"
}

func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressAddressStreetListRequest) SetParamQuery(paramQuery *AddressQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetParamQuery() *AddressQueryDto {
    return r.paramQuery
}

