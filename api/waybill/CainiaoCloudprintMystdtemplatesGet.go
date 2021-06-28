package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
获取用户使用的菜鸟电子面单模板信息 
cainiao.cloudprint.mystdtemplates.get

获取用户使用的菜鸟电子面单
*/
func CainiaoCloudprintMystdtemplatesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintMystdtemplatesGetRequest, session string) (*waybill.CainiaoCloudprintMystdtemplatesGetAPIResponse, error) {
    var resp waybill.CainiaoCloudprintMystdtemplatesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
