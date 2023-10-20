package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpresslogisticsordercreate 快递下单
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
func Alibabaonetouchlogisticsexpresslogisticsordercreate(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpresslogisticsordercreateAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpresslogisticsordercreateAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpresslogisticsordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
