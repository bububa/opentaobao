package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojecttradeitem 新二租品同步
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
func Alibabaalihousenewhomeprojecttradeitem(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojecttradeitemAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojecttradeitemAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojecttradeitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
