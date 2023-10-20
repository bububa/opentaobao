package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressorderdetailget 订单详细信息(面单及仓库信息)
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
func Alibabaonetouchlogisticsexpressorderdetailget(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressorderdetailgetAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressorderdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
