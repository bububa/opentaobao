package damai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damai"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat 
alibaba.damai.mev.open.push.paperformat

pushPaperFormat
*/
func AlibabaDamaiMevOpenPushPaperformat(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushPaperformatRequest, session string) (*damai.AlibabaDamaiMevOpenPushPaperformatAPIResponse, error) {
    var resp damai.AlibabaDamaiMevOpenPushPaperformatAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
