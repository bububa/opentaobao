package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatCardQueryRecord 微信会员卡领取记录查询
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
func AlitripMerchantGalaxyWechatCardQueryRecord(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
