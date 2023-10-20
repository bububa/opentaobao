package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatcardqueryrecord 微信会员卡领取记录查询
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
func Alitripmerchantgalaxywechatcardqueryrecord(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatcardqueryrecordAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatcardqueryrecordAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatcardqueryrecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
