package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTradeAdminCommonOperation 交易迎客松通用服务接口
// yunos.trade.admin.common.operation
//
// 迎客松交易相关通用接口
func YunosTradeAdminCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosTradeAdminCommonOperationAPIRequest, resp *tvupadmin.YunosTradeAdminCommonOperationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
