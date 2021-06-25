package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
获取关键词列表 
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回
*/
func TaobaoWangwangAbstractGetwordlist(clt *core.SDKClient, req *wangwang.TaobaoWangwangAbstractGetwordlistRequest, session string) (*wangwang.TaobaoWangwangAbstractGetwordlistResponse, error) {
    var resp wangwang.TaobaoWangwangAbstractGetwordlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
