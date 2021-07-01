package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
简易商品查询接口 
tmall.nrt.simpleitem.query

为居然之家和阿里的合资公司 homeStyler提供简易的商品信息查询 包含商品名称  图片 状态

后续合资公司服务会迁移到内网 暂时过渡用
*/
func TmallNrtSimpleitemQuery(clt *core.SDKClient, req *nrt.TmallNrtSimpleitemQueryAPIRequest, session string) (*nrt.TmallNrtSimpleitemQueryAPIResponse, error) {
    var resp nrt.TmallNrtSimpleitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
