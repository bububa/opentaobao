package wenyuvideo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wenyuvideo"
)

/* 
根据人物名称查询人物列表 
youku.wenyuvideo.persion.search

根据人物名称查询人物列表
*/
func YoukuWenyuvideoPersionSearch(clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoPersionSearchAPIRequest, session string) (*wenyuvideo.YoukuWenyuvideoPersionSearchAPIResponse, error) {
    var resp wenyuvideo.YoukuWenyuvideoPersionSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
