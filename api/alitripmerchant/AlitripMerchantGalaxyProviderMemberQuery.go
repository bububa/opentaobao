package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyprovidermemberquery 提供会员查询接口
// alitrip.merchant.galaxy.provider.member.query
//
// 星河产品=提供会查询服务
func Alitripmerchantgalaxyprovidermemberquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyprovidermemberqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyprovidermemberqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyprovidermemberqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
