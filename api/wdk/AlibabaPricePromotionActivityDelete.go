package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionActivityDelete 删除档期活动
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
func AlibabaPricePromotionActivityDelete(clt *core.SDKClient, req *wdk.AlibabaPricePromotionActivityDeleteAPIRequest, resp *wdk.AlibabaPricePromotionActivityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
