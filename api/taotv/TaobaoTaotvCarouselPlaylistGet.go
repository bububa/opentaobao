package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvcarouselplaylistget 根据频道ID获取频道下节目单以及当前播放
// taobao.taotv.carousel.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
func Taobaotaotvcarouselplaylistget(clt *core.SDKClient, req *taotv.TaobaotaotvcarouselplaylistgetAPIRequest, session string) (*taotv.TaobaotaotvcarouselplaylistgetAPIResponse, error) {
	var resp taotv.TaobaotaotvcarouselplaylistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
