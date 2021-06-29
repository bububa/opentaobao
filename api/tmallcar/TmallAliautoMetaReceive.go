package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
汽车说明书元数据上传 
tmall.aliauto.meta.receive

天猫汽车对外提供的汽车资源元数据上传接口
*/
func TmallAliautoMetaReceive(clt *core.SDKClient, req *tmallcar.TmallAliautoMetaReceiveRequest, session string) (*tmallcar.TmallAliautoMetaReceiveAPIResponse, error) {
    var resp tmallcar.TmallAliautoMetaReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
