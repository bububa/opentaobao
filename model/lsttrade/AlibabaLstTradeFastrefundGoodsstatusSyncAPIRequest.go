package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest 卖家退款单商品状态同步 API请求
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest struct {
	model.Params
	// 主订单id
	_mainOrderId int64
	// 退款单id
	_refundId string
	// 未发货，枚举类型：UNSEND
	_status string
}

// NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest 初始化AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest对象
func NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest() *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest {
	return &AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.fastrefund.goodsstatus.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetRefundId is RefundId Setter
// 退款单id
func (r *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetStatus is Status Setter
// 未发货，枚举类型：UNSEND
func (r *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetStatus() string {
	return r._status
}
