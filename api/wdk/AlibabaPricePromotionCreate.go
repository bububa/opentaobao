package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionCreate 营销档期活动创建
// alibaba.price.promotion.create
//
// 大润发-盒马帮提供新增创建营销活动
func AlibabaPricePromotionCreate(clt *core.SDKClient, req *wdk.AlibabaPricePromotionCreateAPIRequest, resp *wdk.AlibabaPricePromotionCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
