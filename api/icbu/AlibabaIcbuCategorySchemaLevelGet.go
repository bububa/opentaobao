package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
(新)层级属性获取 
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回
*/
func AlibabaIcbuCategorySchemaLevelGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategorySchemaLevelGetRequest, session string) (*icbu.AlibabaIcbuCategorySchemaLevelGetResponse, error) {
    var resp icbu.AlibabaIcbuCategorySchemaLevelGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
