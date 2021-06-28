package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
（新）ICBU商品发布草稿 
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口
*/
func AlibabaIcbuProductSchemaAddDraft(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaAddDraftRequest, session string) (*icbu.AlibabaIcbuProductSchemaAddDraftAPIResponse, error) {
    var resp icbu.AlibabaIcbuProductSchemaAddDraftAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
