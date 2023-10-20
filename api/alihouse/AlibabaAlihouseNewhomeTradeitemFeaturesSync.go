package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhometradeitemfeaturessync 同步品活动标
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
func Alibabaalihousenewhometradeitemfeaturessync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhometradeitemfeaturessyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhometradeitemfeaturessyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhometradeitemfeaturessyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
