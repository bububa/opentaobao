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
type AlibabaAlihealthMsAreaProvinceListRequest struct {
    model.Params
}

// 初始化AlibabaAlihealthMsAreaProvinceListRequest对象
func NewAlibabaAlihealthMsAreaProvinceListRequest() *AlibabaAlihealthMsAreaProvinceListRequest{
    return &AlibabaAlihealthMsAreaProvinceListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMsAreaProvinceListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.ms.area.province.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMsAreaProvinceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
