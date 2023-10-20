package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelRefundGoodsWaybillAPIRequest) Reset() {
	r._refundWayBillReq = nil
	r.Params.ToZero()
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

var poolAlibabaAscpChannelRefundGoodsWaybillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelRefundGoodsWaybillRequest()
	},
}

// GetAlibabaAscpChannelRefundGoodsWaybillRequest 从 sync.Pool 获取 AlibabaAscpChannelRefundGoodsWaybillAPIRequest
func GetAlibabaAscpChannelRefundGoodsWaybillAPIRequest() *AlibabaAscpChannelRefundGoodsWaybillAPIRequest {
	return poolAlibabaAscpChannelRefundGoodsWaybillAPIRequest.Get().(*AlibabaAscpChannelRefundGoodsWaybillAPIRequest)
}

// ReleaseAlibabaAscpChannelRefundGoodsWaybillAPIRequest 将 AlibabaAscpChannelRefundGoodsWaybillAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelRefundGoodsWaybillAPIRequest(v *AlibabaAscpChannelRefundGoodsWaybillAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelRefundGoodsWaybillAPIRequest.Put(v)
}
