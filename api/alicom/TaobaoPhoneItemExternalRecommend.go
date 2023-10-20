package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneItemExternalRecommend 话费选品能力外放
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
func TaobaoPhoneItemExternalRecommend(clt *core.SDKClient, req *alicom.TaobaoPhoneItemExternalRecommendAPIRequest, resp *alicom.TaobaoPhoneItemExternalRecommendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
