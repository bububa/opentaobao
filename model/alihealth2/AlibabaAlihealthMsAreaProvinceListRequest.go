package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约省份列表查询 API请求
alibaba.alihealth.ms.area.province.list

微信小程序疫苗预约省份列表查询
*/
type AlibabaAlihealthMsAreaProvinceListAPIRequest struct {
    model.Params
}

// 初始化AlibabaAlihealthMsAreaProvinceListAPIRequest对象
func NewAlibabaAlihealthMsAreaProvinceListRequest() *AlibabaAlihealthMsAreaProvinceListAPIRequest{
    return &AlibabaAlihealthMsAreaProvinceListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMsAreaProvinceListAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.ms.area.province.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMsAreaProvinceListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
