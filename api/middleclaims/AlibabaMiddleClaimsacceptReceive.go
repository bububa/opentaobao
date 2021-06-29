package middleclaims

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/middleclaims"
)

/* 
国际化中台服务域接收保险公司理赔受理结果 
alibaba.middle.claimsaccept.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
*/
func AlibabaMiddleClaimsacceptReceive(clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsacceptReceiveRequest, session string) (*middleclaims.AlibabaMiddleClaimsacceptReceiveAPIResponse, error) {
    var resp middleclaims.AlibabaMiddleClaimsacceptReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
