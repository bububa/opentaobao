package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
模糊查询服务初始化 
taobao.wangwang.abstract.initialize

模糊查询服务初始化，只支持json返回
*/
func TaobaoWangwangAbstractInitialize(clt *core.SDKClient, req *wangwang.TaobaoWangwangAbstractInitializeRequest, session string) (*wangwang.TaobaoWangwangAbstractInitializeResponse, error) {
    var resp wangwang.TaobaoWangwangAbstractInitializeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
