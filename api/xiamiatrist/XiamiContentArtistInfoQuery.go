package xiamiatrist

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamiatrist"
)

/* 
搜索艺人列表 
xiami.content.artist.info.query

根据查询条件，搜索相关艺人列表
*/
func XiamiContentArtistInfoQuery(clt *core.SDKClient, req *xiamiatrist.XiamiContentArtistInfoQueryRequest, session string) (*xiamiatrist.XiamiContentArtistInfoQueryAPIResponse, error) {
    var resp xiamiatrist.XiamiContentArtistInfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
