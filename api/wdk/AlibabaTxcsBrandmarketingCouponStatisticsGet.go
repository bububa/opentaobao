package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatxcsbrandmarketingcouponstatisticsget 品牌营销导购员券推广统计数据回流
// alibaba.txcs.brandmarketing.coupon.statistics.get
//
// 请求券统计数据回流
func Alibabatxcsbrandmarketingcouponstatisticsget(clt *core.SDKClient, req *wdk.AlibabatxcsbrandmarketingcouponstatisticsgetAPIRequest, session string) (*wdk.AlibabatxcsbrandmarketingcouponstatisticsgetAPIResponse, error) {
	var resp wdk.AlibabatxcsbrandmarketingcouponstatisticsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
