package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatcardparmquery 微信会员卡添加
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
func Alitripmerchantgalaxywechatcardparmquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatcardparmqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatcardparmqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatcardparmqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
