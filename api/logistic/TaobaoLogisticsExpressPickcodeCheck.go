package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressPickcodeCheck 快递公司取货码校验
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
func TaobaoLogisticsExpressPickcodeCheck(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressPickcodeCheckAPIRequest, session string) (*logistic.TaobaoLogisticsExpressPickcodeCheckAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressPickcodeCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
