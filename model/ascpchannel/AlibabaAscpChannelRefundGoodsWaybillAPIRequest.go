package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundGoodsWaybillAPIRequest 淘外分销退货回传物流单号 API请求
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
type AlibabaAscpChannelRefundGoodsWaybillAPIRequest struct {
	model.Params
	// 请求
	_refundWayBillReq *ExternalRefundGoodsWaybillRequest
}

// NewAlibabaAscpChannelRefundGoodsWaybillRequest 初始化AlibabaAscpChannelRefundGoodsWaybillAPIRequest对象
func NewAlibabaAscpChannelRefundGoodsWaybillRequest() *AlibabaAscpChannelRefundGoodsWaybillAPIRequest {
	return &AlibabaAscpChannelRefundGoodsWaybillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelRefundGoodsWaybillAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.goods.waybill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelRefundGoodsWaybillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelRefundGoodsWaybillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundWayBillReq is RefundWayBillReq Setter
// 请求
func (r *AlibabaAscpChannelRefundGoodsWaybillAPIRequest) SetRefundWayBillReq(_refundWayBillReq *ExternalRefundGoodsWaybillRequest) error {
	r._refundWayBillReq = _refundWayBillReq
	r.Set("refund_way_bill_req", _refundWayBillReq)
	return nil
}

// GetRefundWayBillReq RefundWayBillReq Getter
func (r AlibabaAscpChannelRefundGoodsWaybillAPIRequest) GetRefundWayBillReq() *ExternalRefundGoodsWaybillRequest {
	return r._refundWayBillReq
}
