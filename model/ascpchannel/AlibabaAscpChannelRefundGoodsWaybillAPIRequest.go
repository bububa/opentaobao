package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelrefundgoodswaybillAPIRequest 淘外分销退货回传物流单号 API请求
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
type AlibabaascpchannelrefundgoodswaybillAPIRequest struct {
	model.Params
	// 请求
	_refundWayBillReq *ExternalRefundGoodsWaybillRequest
}

// NewAlibabaascpchannelrefundgoodswaybillRequest 初始化AlibabaascpchannelrefundgoodswaybillAPIRequest对象
func NewAlibabaascpchannelrefundgoodswaybillRequest() *AlibabaascpchannelrefundgoodswaybillAPIRequest {
	return &AlibabaascpchannelrefundgoodswaybillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelrefundgoodswaybillAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.goods.waybill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelrefundgoodswaybillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelrefundgoodswaybillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundWayBillReq is RefundWayBillReq Setter
// 请求
func (r *AlibabaascpchannelrefundgoodswaybillAPIRequest) SetRefundWayBillReq(_refundWayBillReq *ExternalRefundGoodsWaybillRequest) error {
	r._refundWayBillReq = _refundWayBillReq
	r.Set("refund_way_bill_req", _refundWayBillReq)
	return nil
}

// GetRefundWayBillReq RefundWayBillReq Getter
func (r AlibabaascpchannelrefundgoodswaybillAPIRequest) GetRefundWayBillReq() *ExternalRefundGoodsWaybillRequest {
	return r._refundWayBillReq
}
