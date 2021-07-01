package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelRefundGoodsWaybillAPIRequest
淘外分销退货回传物流单号 API请求
alibaba.ascp.channel.refund.goods.waybill

淘外分销退货回传物流单号 */
type AlibabaAscpChannelRefundGoodsWaybillAPIRequest struct {
	model.Params
	// 请求
	_refundWayBillReq *ExternalRefundGoodsWaybillRequest
}

// New
