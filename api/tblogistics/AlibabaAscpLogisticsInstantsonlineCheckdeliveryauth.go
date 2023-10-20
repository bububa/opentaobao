package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineCheckdeliveryauth 同城配送在线下单检查授权
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
func AlibabaAscpLogisticsInstantsonlineCheckdeliveryauth(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
