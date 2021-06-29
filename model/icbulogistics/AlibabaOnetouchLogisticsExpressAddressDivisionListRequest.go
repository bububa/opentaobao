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
type AlibabaOnetouchLogisticsExpressAddressDivisionListRequest struct {
    model.Params
    // 请求参数
    paramQuery   *AddressQueryDto
}

// 初始化AlibabaOnetouchLogisticsExpressAddressDivisionListRequest对象
func NewAlibabaOnetouchLogisticsExpressAddressDivisionListRequest() *AlibabaOnetouchLogisticsExpressAddressDivisionListRequest{
    return &AlibabaOnetouchLogisticsExpressAddressDivisionListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.address.division.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) SetParamQuery(paramQuery *AddressQueryDto) error {
    r.paramQuery = paramQuery
    r.Set("param_query", paramQuery)
    return nil
}

// ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressAddressDivisionListRequest) GetParamQuery() *AddressQueryDto {
    return r.paramQuery
}
