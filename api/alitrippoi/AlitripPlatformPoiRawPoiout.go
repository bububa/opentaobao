package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// AlitripPlatformPoiRawPoiout 飞猪poi输出
// alitrip.platform.poi.raw.poiout
//
// 输出指定城市poi指定信息
func AlitripPlatformPoiRawPoiout(clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawPoioutAPIRequest, resp *alitrippoi.AlitripPlatformPoiRawPoioutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
