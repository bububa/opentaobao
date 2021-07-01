package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
绑定物理卡 
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起
*/
func AlibabaAlscCrmCardBindcard(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBindcardAPIRequest, session string) (*alsc.AlibabaAlscCrmCardBindcardAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardBindcardAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
