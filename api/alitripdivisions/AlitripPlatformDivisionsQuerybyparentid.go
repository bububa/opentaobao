package alitripdivisions

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripdivisions"
)

// AlitripPlatformDivisionsQuerybyparentid 根据父节点id查询下级行政区划数据
// alitrip.platform.divisions.querybyparentid
//
// 根据行政区划id查询下一层级行政区划数据
func AlitripPlatformDivisionsQuerybyparentid(ctx context.Context, clt *core.SDKClient, req *alitripdivisions.AlitripPlatformDivisionsQuerybyparentidAPIRequest, resp *alitripdivisions.AlitripPlatformDivisionsQuerybyparentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
