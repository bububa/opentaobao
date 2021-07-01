package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品编辑获取schema信息 
alibaba.item.edit.schema.get

商品编辑时，获取商品规则信息
*/
func AlibabaItemEditSchemaGet(clt *core.SDKClient, req *product.AlibabaItemEditSchemaGetAPIRequest, session string) (*product.AlibabaItemEditSchemaGetAPIResponse, error) {
    var resp product.AlibabaItemEditSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
