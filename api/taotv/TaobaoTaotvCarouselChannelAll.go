package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
获取所有频道列表 
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序
*/
func TaobaoTaotvCarouselChannelAll(clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselChannelAllAPIRequest, session string) (*taotv.TaobaoTaotvCarouselChannelAllAPIResponse, error) {
    var resp taotv.TaobaoTaotvCarouselChannelAllAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
