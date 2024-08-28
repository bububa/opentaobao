package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripTravelGereralproductUpdate 通用类目产品发布编辑
// alitrip.travel.gereralproduct.update
//
// 提供给飞猪供销平台供应商发布编辑通用类目产品的API
func AlitripTravelGereralproductUpdate(ctx context.Context, clt *core.SDKClient, req *travel.AlitripTravelGereralproductUpdateAPIRequest, resp *travel.AlitripTravelGereralproductUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
