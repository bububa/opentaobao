package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
获取商家单一自定义区 
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区
*/
func CainiaoCloudprintSingleCustomareaGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintSingleCustomareaGetRequest, session string) (*waybill.CainiaoCloudprintSingleCustomareaGetResponse, error) {
    var resp waybill.CainiaoCloudprintSingleCustomareaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
