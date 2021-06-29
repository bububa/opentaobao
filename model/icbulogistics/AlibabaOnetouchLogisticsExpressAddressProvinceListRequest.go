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
type AlibabaOnetouchLogisticsExpressAddressProvinceListRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDto
}

// 初始化AlibabaOnetouchLogisticsExpressAddressProvinceListRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressProvinceListRequest() *AlibabaOnetouchLogisticsExpressAddressProvinceListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressProvinceListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.province.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressProvinceListRequest) GetParamQuery() *AddressQueryDto {
    return r._paramQuery
}