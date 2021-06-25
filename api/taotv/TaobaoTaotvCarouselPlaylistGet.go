package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
根据频道ID获取频道下节目单以及当前播放 
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
*/
func TaobaoTaotvCarouselPlaylistGet(clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselPlaylistGetRequest, session string) (*taotv.TaobaoTaotvCarouselPlaylistGetResponse, error) {
    var resp taotv.TaobaoTaotvCarouselPlaylistGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
