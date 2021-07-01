package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaPricePromotionActivityDelete
删除档期活动
alibaba.price.promotion.activity.delete

删除盒马帮档期活动 */
func AlibabaPricePromotionActivityDelete(clt *core.SDKClient, req *wdk.AlibabaPricePromotionActivityDeleteAPIRequest, session string) (*wdk.AlibabaPricePromotionActivityDeleteAPIResponse, error) {
	var resp wdk.AlibabaPricePromotionActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
