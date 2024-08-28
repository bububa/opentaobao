package alitrippoi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// AlitripPlatformPoiRawPoioutbypoiids 根据poiId输出飞猪poi数据
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
func AlitripPlatformPoiRawPoioutbypoiids(ctx context.Context, clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawPoioutbypoiidsAPIRequest, resp *alitrippoi.AlitripPlatformPoiRawPoioutbypoiidsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
