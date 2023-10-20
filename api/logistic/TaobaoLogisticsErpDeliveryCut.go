package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsErpDeliveryCut ERP发起配拦截
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
func TaobaoLogisticsErpDeliveryCut(clt *core.SDKClient, req *logistic.TaobaoLogisticsErpDeliveryCutAPIRequest, session string) (*logistic.TaobaoLogisticsErpDeliveryCutAPIResponse, error) {
	var resp logistic.TaobaoLogisticsErpDeliveryCutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
