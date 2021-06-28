package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
ICBU商品发布草稿接口 
alibaba.icbu.product.add.draft

发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
*/
func AlibabaIcbuProductAddDraft(clt *core.SDKClient, req *icbu.AlibabaIcbuProductAddDraftRequest, session string) (*icbu.AlibabaIcbuProductAddDraftAPIResponse, error) {
    var resp icbu.AlibabaIcbuProductAddDraftAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
