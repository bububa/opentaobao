package crm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/crm"
)

/* 
删除人群实例中的指定买家 
taobao.crm.exchange.crowdinstance.delete

删除人群实例中的指定买家
*/
func TaobaoCrmExchangeCrowdinstanceDelete(clt *core.SDKClient, req *crm.TaobaoCrmExchangeCrowdinstanceDeleteRequest, session string) (*crm.TaobaoCrmExchangeCrowdinstanceDeleteResponse, error) {
    var resp crm.TaobaoCrmExchangeCrowdinstanceDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
