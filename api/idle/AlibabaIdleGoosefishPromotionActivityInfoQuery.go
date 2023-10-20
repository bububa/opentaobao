package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleGoosefishPromotionActivityInfoQuery 闲鱼三方活动参与信息查询
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
func AlibabaIdleGoosefishPromotionActivityInfoQuery(clt *core.SDKClient, req *idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest, session string) (*idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIResponse, error) {
	var resp idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
