package alidoc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alidoc"
)

/* 
更新药店 
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口
*/
func AlibabaAlihealthAlidocDrugStoreUpdate(clt *core.SDKClient, req *alidoc.AlibabaAlihealthAlidocDrugStoreUpdateRequest, session string) (*alidoc.AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse, error) {
    var resp alidoc.AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}