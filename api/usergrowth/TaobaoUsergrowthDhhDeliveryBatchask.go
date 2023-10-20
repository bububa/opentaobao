package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoUsergrowthDhhDeliveryBatchask 广告曝光前判定批量接口V2
// taobao.usergrowth.dhh.delivery.batchask
//
// 广告曝光前判定批量接口V2
func TaobaoUsergrowthDhhDeliveryBatchask(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest, resp *usergrowth.TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
