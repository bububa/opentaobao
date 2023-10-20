package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradefastrefundgoodsstatussyncAPIRequest 卖家退款单商品状态同步 API请求
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
type AlibabalsttradefastrefundgoodsstatussyncAPIRequest struct {
	model.Params
	// 退款单id
	_refundId string
	// 未发货，枚举类型：UNSEND
	_status string
	// 主订单id
	_mainOrderId int64
}

// NewAlibabalsttradefastrefundgoodsstatussyncRequest 初始化AlibabalsttradefastrefundgoodsstatussyncAPIRequest对象
func NewAlibabalsttradefastrefundgoodsstatussyncRequest() *AlibabalsttradefastrefundgoodsstatussyncAPIRequest {
	return &AlibabalsttradefastrefundgoodsstatussyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.fastrefund.goodsstatus.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款单id
func (r *AlibabalsttradefastrefundgoodsstatussyncAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetStatus is Status Setter
// 未发货，枚举类型：UNSEND
func (r *AlibabalsttradefastrefundgoodsstatussyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetStatus() string {
	return r._status
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabalsttradefastrefundgoodsstatussyncAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabalsttradefastrefundgoodsstatussyncAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
