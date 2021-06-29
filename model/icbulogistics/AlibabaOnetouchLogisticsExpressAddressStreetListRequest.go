package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-街道 API请求
alibaba.onetouch.logistics.express.address.street.list

四级地址库-街道模糊查询
*/
type AlibabaOnetouchLogisticsExpressAddressStreetListRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDTO
}

// 初始化AlibabaOnetouchLogisticsExpressAddressStreetListRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressStreetListRequest() *AlibabaOnetouchLogisticsExpressAddressStreetListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressStreetListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.street.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressStreetListRequest) SetParamQuery(_paramQuery *AddressQueryDTO) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressStreetListRequest) GetParamQuery() *AddressQueryDTO {
    return r._paramQuery
}
