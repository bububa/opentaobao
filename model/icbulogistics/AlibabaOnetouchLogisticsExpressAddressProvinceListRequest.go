package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-省 APIRequest
alibaba.onetouch.logistics.express.address.province.list

四级地址库-省
*/
type AlibabaOnetouchLogisticsExpressAddressProvinceListRequest struct {
    model.Params

    // 请求参数
    paramQuery   *AddressQueryDto 

}

func NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest() *AlibabaOnetouchLogisticsExpressAddressProvinceListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressProvinceListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.province.list"
}

func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) SetParamQuery(paramQuery *AddressQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetParamQuery() *AddressQueryDto {
    return r.paramQuery
}

