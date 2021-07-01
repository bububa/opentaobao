package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-省 API请求
alibaba.onetouch.logistics.express.address.province.list

四级地址库-省
*/
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDTO
}

// 初始化AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest() *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest{
    return &AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.province.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDTO) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest) GetParamQuery() *AddressQueryDTO {
    return r._paramQuery
}
