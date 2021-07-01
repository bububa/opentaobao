package xiamicontent

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamicontent"
)

/* 
获取歌单详情接口 
xiami.content.songs.collect.get

根据歌单id，获取歌单详情
*/
func XiamiContentSongsCollectGet(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsCollectGetAPIRequest, session string) (*xiamicontent.XiamiContentSongsCollectGetAPIResponse, error) {
    var resp xiamicontent.XiamiContentSongsCollectGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
