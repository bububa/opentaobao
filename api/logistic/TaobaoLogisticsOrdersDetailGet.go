package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoLogisticsOrdersDetailGet
批量查询物流订单,返回详细信息
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。 */
func TaobaoLogisticsOrdersDetailGet(clt *core.SDKClient, req *logistic.TaobaoLogisticsOrdersDetailGetAPIRequest, session string) (*logistic.TaobaoLogisticsOrdersDetailGetAPIResponse, error) {
	var resp logistic.TaobaoLogisticsOrdersDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
