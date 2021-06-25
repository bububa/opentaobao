package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-物料搜索 
taobao.tbk.dg.material.optional

通用物料搜索API（导购）
*/
func TaobaoTbkDgMaterialOptional(clt *core.SDKClient, req *tbk.TaobaoTbkDgMaterialOptionalRequest, session string) (*tbk.TaobaoTbkDgMaterialOptionalResponse, error) {
    var resp tbk.TaobaoTbkDgMaterialOptionalAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
