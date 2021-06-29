package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约省份列表查询 APIRequest
alibaba.alihealth.ms.area.province.list

微信小程序疫苗预约省份列表查询
*/
type AlibabaAlihealthMsAreaProvinceListRequest struct {
    model.Params

}

func NewAlibabaAlihealthMsAreaProvinceListRequest() *AlibabaAlihealthMsAreaProvinceListRequest{
    return &AlibabaAlihealthMsAreaProvinceListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMsAreaProvinceListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.ms.area.province.list"
}

func (r AlibabaAlihealthMsAreaProvinceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


