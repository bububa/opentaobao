package middleclaims

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/middleclaims"
)

/* 
国际化中台服务域接收理赔结果 
alibaba.middle.claimsresult.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
*/
func AlibabaMiddleClaimsresultReceive(clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsresultReceiveAPIRequest, session string) (*middleclaims.AlibabaMiddleClaimsresultReceiveAPIResponse, error) {
    var resp middleclaims.AlibabaMiddleClaimsresultReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
