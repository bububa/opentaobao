package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest 卖家退款单商品状态同步 API请求
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest struct {
	model.Params
	// 退款单id
	_refundId string
	// 未发货，枚举类型：UNSEND
	_status string
	// 主订单id
	_mainOrderId int64
}

// NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest 初始化AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest对象
func NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest() *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest {
	return &AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) Reset() {
	r._refundId = ""
	r._status = ""
	r._mainOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.fastrefund.goodsstatus.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeFastrefundGoodsstatusSyncRequest()
	},
}

// GetAlibabaLstTradeFastrefundGoodsstatusSyncRequest 从 sync.Pool 获取 AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest
func GetAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest() *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest {
	return poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest.Get().(*AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest)
}

// ReleaseAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest 将 AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest(v *AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest.Put(v)
}
