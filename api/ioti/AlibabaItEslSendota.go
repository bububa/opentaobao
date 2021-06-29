package ioti

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ioti"
)

/* 
电子价签ota接口 
alibaba.it.esl.sendota

厂测接口，电子价签ota接口
*/
func AlibabaItEslSendota(clt *core.SDKClient, req *ioti.AlibabaItEslSendotaRequest, session string) (*ioti.AlibabaItEslSendotaAPIResponse, error) {
    var resp ioti.AlibabaItEslSendotaAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
