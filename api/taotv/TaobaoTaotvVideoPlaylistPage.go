package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
分页获取所有播单 
taobao.taotv.video.playlist.page

获取所有播单信息（分页）
*/
func TaobaoTaotvVideoPlaylistPage(clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistPageRequest, session string) (*taotv.TaobaoTaotvVideoPlaylistPageResponse, error) {
    var resp taotv.TaobaoTaotvVideoPlaylistPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
