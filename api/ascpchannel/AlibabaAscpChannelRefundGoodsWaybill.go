package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelRefundGoodsWaybill 淘外分销退货回传物流单号
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
func AlibabaAscpChannelRefundGoodsWaybill(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelRefundGoodsWaybillAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelRefundGoodsWaybillAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelRefundGoodsWaybillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
