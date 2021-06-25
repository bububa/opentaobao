package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
删除运费模板 
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
func TaobaoDeliveryTemplateDelete(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplateDeleteRequest, session string) (*logistic.TaobaoDeliveryTemplateDeleteResponse, error) {
    var resp logistic.TaobaoDeliveryTemplateDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
