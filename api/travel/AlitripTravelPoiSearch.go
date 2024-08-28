package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripTravelPoiSearch POI信息查询
// alitrip.travel.poi.search
//
// POI信息查询，用于商品更新使用
func AlitripTravelPoiSearch(ctx context.Context, clt *core.SDKClient, req *travel.AlitripTravelPoiSearchAPIRequest, resp *travel.AlitripTravelPoiSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
