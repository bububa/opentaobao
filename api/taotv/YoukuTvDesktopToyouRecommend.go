package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// YoukuTvDesktopToyouRecommend TV桌面为你推荐接口
// youku.tv.desktop.toyou.recommend
//
// 提供为你推荐数据
func YoukuTvDesktopToyouRecommend(clt *core.SDKClient, req *taotv.YoukuTvDesktopToyouRecommendAPIRequest, resp *taotv.YoukuTvDesktopToyouRecommendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
