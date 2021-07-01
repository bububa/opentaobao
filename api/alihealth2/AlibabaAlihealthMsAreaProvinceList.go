package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
疫苗预约省份列表查询 
alibaba.alihealth.ms.area.province.list

微信小程序疫苗预约省份列表查询
*/
func AlibabaAlihealthMsAreaProvinceList(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMsAreaProvinceListAPIRequest, session string) (*alihealth2.AlibabaAlihealthMsAreaProvinceListAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthMsAreaProvinceListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
