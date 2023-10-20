package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCommonGetEnumsbyname 小程序公共枚举类获取接口
// alitrip.merchant.galaxy.common.get.enumsbyname
//
// 通过枚举名称获取枚举类实例内容
func AlitripMerchantGalaxyCommonGetEnumsbyname(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCommonGetEnumsbynameAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
