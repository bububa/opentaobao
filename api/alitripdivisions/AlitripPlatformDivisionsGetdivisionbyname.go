package alitripdivisions

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripdivisions"
)

// AlitripPlatformDivisionsGetdivisionbyname 根据中文名称与行政区划级别查询行政区划数据
// alitrip.platform.divisions.getdivisionbyname
//
// 根据中文名称与行政区划级别查询行政区划数据
func AlitripPlatformDivisionsGetdivisionbyname(ctx context.Context, clt *core.SDKClient, req *alitripdivisions.AlitripPlatformDivisionsGetdivisionbynameAPIRequest, resp *alitripdivisions.AlitripPlatformDivisionsGetdivisionbynameAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
