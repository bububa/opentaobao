package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
获取用户下所有模板 
taobao.delivery.templates.get

根据用户ID获取用户下所有模板
*/
func TaobaoDeliveryTemplatesGet(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplatesGetRequest, session string) (*logistic.TaobaoDeliveryTemplatesGetAPIResponse, error) {
    var resp logistic.TaobaoDeliveryTemplatesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
