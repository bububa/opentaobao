package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpressorderpaytmsquery 上门取退运费支付状态查询接口
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
func Taobaologisticsexpressorderpaytmsquery(clt *core.SDKClient, req *logistic.TaobaologisticsexpressorderpaytmsqueryAPIRequest, session string) (*logistic.TaobaologisticsexpressorderpaytmsqueryAPIResponse, error) {
	var resp logistic.TaobaologisticsexpressorderpaytmsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
