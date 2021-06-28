package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询物理卡 
alibaba.alsc.crm.card.qryphysical

查询物理卡
*/
func AlibabaAlscCrmCardQryphysical(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQryphysicalRequest, session string) (*alsc.AlibabaAlscCrmCardQryphysicalAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardQryphysicalAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
