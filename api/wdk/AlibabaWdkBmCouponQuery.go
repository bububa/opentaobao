package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmCouponQuery 淘鲜达券信息查询接口
// alibaba.wdk.bm.coupon.query
//
// 淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
func AlibabaWdkBmCouponQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmCouponQueryAPIRequest, resp *wdk.AlibabaWdkBmCouponQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
