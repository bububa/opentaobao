package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatinfo 星河-获取微信用户的信息
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
func Alitripmerchantgalaxywechatinfo(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatinfoAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatinfoAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
