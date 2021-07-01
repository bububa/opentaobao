package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
isv资源查询 
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等
*/
func CainiaoCloudprintIsvResourcesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintIsvResourcesGetAPIRequest, session string) (*waybill.CainiaoCloudprintIsvResourcesGetAPIResponse, error) {
    var resp waybill.CainiaoCloudprintIsvResourcesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
