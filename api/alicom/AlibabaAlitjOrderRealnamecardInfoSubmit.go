package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
阿里实人认证卡片信息回传 
alibaba.alitj.order.realnamecard.info.submit

阿里实人认证卡片信息回传。ISP相关商家在线对接阿里通信的实人认证功能，在线提交订单对应运营商的合约订购相关信息，以便完成在线使用实人认证功能。
*/
func AlibabaAlitjOrderRealnamecardInfoSubmit(clt *core.SDKClient, req *alicom.AlibabaAlitjOrderRealnamecardInfoSubmitRequest, session string) (*alicom.AlibabaAlitjOrderRealnamecardInfoSubmitAPIResponse, error) {
    var resp alicom.AlibabaAlitjOrderRealnamecardInfoSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
