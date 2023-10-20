package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvcarouselchannelall 获取所有频道列表
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
func Taobaotaotvcarouselchannelall(clt *core.SDKClient, req *taotv.TaobaotaotvcarouselchannelallAPIRequest, session string) (*taotv.TaobaotaotvcarouselchannelallAPIResponse, error) {
	var resp taotv.TaobaotaotvcarouselchannelallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
