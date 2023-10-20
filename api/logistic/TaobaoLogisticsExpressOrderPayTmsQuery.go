package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressOrderPayTmsQuery 上门取退运费支付状态查询接口
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
func TaobaoLogisticsExpressOrderPayTmsQuery(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIRequest, session string) (*logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
