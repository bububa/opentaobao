package vaccin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/vaccin"
)

/* 
疫苗预约地市信息查询 
alibaba.alihealth.ms.area.district.list

微信小程序疫苗预约地市信息查询
*/
func AlibabaAlihealthMsAreaDistrictList(clt *core.SDKClient, req *vaccin.AlibabaAlihealthMsAreaDistrictListAPIRequest, session string) (*vaccin.AlibabaAlihealthMsAreaDistrictListAPIResponse, error) {
    var resp vaccin.AlibabaAlihealthMsAreaDistrictListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
