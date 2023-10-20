package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhometradeitemrelation 货独立绑定货品
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
func Alibabaalihousenewhometradeitemrelation(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhometradeitemrelationAPIRequest, session string) (*alihouse.AlibabaalihousenewhometradeitemrelationAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhometradeitemrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
