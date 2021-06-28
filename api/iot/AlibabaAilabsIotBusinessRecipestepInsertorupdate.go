package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
插入或更新食谱步骤 
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤
*/
func AlibabaAilabsIotBusinessRecipestepInsertorupdate(clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateRequest, session string) (*iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse, error) {
    var resp iot.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
