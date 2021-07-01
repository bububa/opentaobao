package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
删除关键词 
taobao.wangwang.abstract.deleteword

删除关键词，只支持json返回
*/
func TaobaoWangwangAbstractDeleteword(clt *core.SDKClient, req *wangwang.TaobaoWangwangAbstractDeletewordAPIRequest, session string) (*wangwang.TaobaoWangwangAbstractDeletewordAPIResponse, error) {
    var resp wangwang.TaobaoWangwangAbstractDeletewordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
