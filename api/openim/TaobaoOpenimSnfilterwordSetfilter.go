package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
关键词过滤 
taobao.openim.snfilterword.setfilter

设置openim关键词过滤
*/
func TaobaoOpenimSnfilterwordSetfilter(clt *core.SDKClient, req *openim.TaobaoOpenimSnfilterwordSetfilterRequest, session string) (*openim.TaobaoOpenimSnfilterwordSetfilterResponse, error) {
    var resp openim.TaobaoOpenimSnfilterwordSetfilterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
