package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
（新）渲染草稿商品数据 
alibaba.icbu.product.schema.render.draft

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
*/
func AlibabaIcbuProductSchemaRenderDraft(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaRenderDraftRequest, session string) (*icbu.AlibabaIcbuProductSchemaRenderDraftResponse, error) {
    var resp icbu.AlibabaIcbuProductSchemaRenderDraftAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
