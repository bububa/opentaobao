package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// AlitripPlatformPoiRawFeed 存储poi原始数据
// alitrip.platform.poi.raw.feed
//
// 对接外部数据源，外部数据推送poi数据到飞猪
func AlitripPlatformPoiRawFeed(clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawFeedAPIRequest, resp *alitrippoi.AlitripPlatformPoiRawFeedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
