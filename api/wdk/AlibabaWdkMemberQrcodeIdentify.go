package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
扫码识别会员接口 
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
*/
func AlibabaWdkMemberQrcodeIdentify(clt *core.SDKClient, req *wdk.AlibabaWdkMemberQrcodeIdentifyRequest, session string) (*wdk.AlibabaWdkMemberQrcodeIdentifyAPIResponse, error) {
    var resp wdk.AlibabaWdkMemberQrcodeIdentifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
