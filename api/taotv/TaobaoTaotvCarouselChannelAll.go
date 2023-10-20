package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// TaobaoTaotvCarouselChannelAll 获取所有频道列表
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
func TaobaoTaotvCarouselChannelAll(clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselChannelAllAPIRequest, resp *taotv.TaobaoTaotvCarouselChannelAllAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
