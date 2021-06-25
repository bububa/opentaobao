package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
修改运费模板 
taobao.delivery.template.update

修改运费模板
*/
func TaobaoDeliveryTemplateUpdate(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplateUpdateRequest, session string) (*logistic.TaobaoDeliveryTemplateUpdateResponse, error) {
    var resp logistic.TaobaoDeliveryTemplateUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
