package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressPickcodeCheck 快递公司取货码校验
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
func TaobaoLogisticsExpressPickcodeCheck(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressPickcodeCheckAPIRequest, resp *logistic.TaobaoLogisticsExpressPickcodeCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
