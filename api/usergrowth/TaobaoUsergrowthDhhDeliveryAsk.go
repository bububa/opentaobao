package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoUsergrowthDhhDeliveryAsk 广告曝光前判定接口V2
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
func TaobaoUsergrowthDhhDeliveryAsk(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIRequest, resp *usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
