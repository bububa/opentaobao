package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasethirdordersync 第三方订单同步
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
func Alibabaalihealthmedicalbasethirdordersync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasethirdordersyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasethirdordersyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasethirdordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
