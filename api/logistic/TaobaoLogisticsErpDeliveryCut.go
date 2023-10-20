package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticserpdeliverycut ERP发起配拦截
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
func Taobaologisticserpdeliverycut(clt *core.SDKClient, req *logistic.TaobaologisticserpdeliverycutAPIRequest, session string) (*logistic.TaobaologisticserpdeliverycutAPIResponse, error) {
	var resp logistic.TaobaologisticserpdeliverycutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
