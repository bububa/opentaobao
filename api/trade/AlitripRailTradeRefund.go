package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alitriprailtraderefund 退票接口
// alitrip.rail.trade.refund
//
// 退票接口
func Alitriprailtraderefund(clt *core.SDKClient, req *trade.AlitriprailtraderefundAPIRequest, session string) (*trade.AlitriprailtraderefundAPIResponse, error) {
	var resp trade.AlitriprailtraderefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
