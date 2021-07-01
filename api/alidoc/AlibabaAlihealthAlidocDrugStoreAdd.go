package alidoc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alidoc"
)

/* 
gsk新增药店 
alibaba.alihealth.alidoc.drug.store.add

GSK上传药店信息
*/
func AlibabaAlihealthAlidocDrugStoreAdd(clt *core.SDKClient, req *alidoc.AlibabaAlihealthAlidocDrugStoreAddAPIRequest, session string) (*alidoc.AlibabaAlihealthAlidocDrugStoreAddAPIResponse, error) {
    var resp alidoc.AlibabaAlihealthAlidocDrugStoreAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
