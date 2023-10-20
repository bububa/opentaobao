package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionActivityQuery 查询盒马帮档期活动详情
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
func AlibabaPricePromotionActivityQuery(clt *core.SDKClient, req *wdk.AlibabaPricePromotionActivityQueryAPIRequest, session string) (*wdk.AlibabaPricePromotionActivityQueryAPIResponse, error) {
	var resp wdk.AlibabaPricePromotionActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
