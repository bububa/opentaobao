package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-市 API请求
alibaba.onetouch.logistics.express.address.city.list

四级地址库-市
*/
type AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDTO
}

// 初始化AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressCityListRequest() *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest{
    return &AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.city.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDTO) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest) GetParamQuery() *AddressQueryDTO {
    return r._paramQuery
}
