package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
AG物流签收状态写接口 
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
func TaobaoNextoneLogisticsSignUpdate(clt *core.SDKClient, req *logistic.TaobaoNextoneLogisticsSignUpdateRequest, session string) (*logistic.TaobaoNextoneLogisticsSignUpdateResponse, error) {
    var resp logistic.TaobaoNextoneLogisticsSignUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
