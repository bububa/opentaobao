package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
ott播单 
taobao.taotv.video.playlist.ottnav.get

根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
*/
func TaobaoTaotvVideoPlaylistOttnavGet(clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistOttnavGetRequest, session string) (*taotv.TaobaoTaotvVideoPlaylistOttnavGetResponse, error) {
    var resp taotv.TaobaoTaotvVideoPlaylistOttnavGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
