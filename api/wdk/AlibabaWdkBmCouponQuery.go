package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbmcouponquery 淘鲜达券信息查询接口
// alibaba.wdk.bm.coupon.query
//
// 淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
func Alibabawdkbmcouponquery(clt *core.SDKClient, req *wdk.AlibabawdkbmcouponqueryAPIRequest, session string) (*wdk.AlibabawdkbmcouponqueryAPIResponse, error) {
	var resp wdk.AlibabawdkbmcouponqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
