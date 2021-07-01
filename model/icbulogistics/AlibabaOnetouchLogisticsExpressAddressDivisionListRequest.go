package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-区域 API请求
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区
*/
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest struct {
    model.Params
    // 请求参数
    _paramQuery   *AddressQueryDTO
}

// 初始化AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest{
    return &AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.division.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) SetParamQuery(_paramQuery *AddressQueryDTO) error {
    r._paramQuery = _paramQuery
    r.Set("param_query", _paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest) GetParamQuery() *AddressQueryDTO {
    return r._paramQuery
}
