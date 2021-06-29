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
type AlibabaOnetouchLogisticsExpressAddressCityListRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDto
}

// 初始化AlibabaOnetouchLogisticsExpressAddressCityListRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressCityListRequest() *AlibabaOnetouchLogisticsExpressAddressCityListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressCityListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.city.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressCityListRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressCityListRequest) GetParamQuery() *AddressQueryDto {
    return r._paramQuery
}
