package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest 供应商/分销商关闭采购申请/经销采购单 API请求
// taobao.fenxiao.dealer.requisitionorder.close
//
// 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest struct {
	model.Params
	// 关闭详细原因，字数5-200字
	_reasonDetail string
	// 经销采购单编号
	_dealerOrderId int64
	// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
	_reason int64
}

// NewTaobaoFenxiaoDealerRequisitionorderCloseRequest 初始化TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderCloseRequest() *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) Reset() {
	r._reasonDetail = ""
	r._dealerOrderId = 0
	r._reason = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReasonDetail is ReasonDetail Setter
// 关闭详细原因，字数5-200字
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetReasonDetail(_reasonDetail string) error {
	r._reasonDetail = _reasonDetail
	r.Set("reason_detail", _reasonDetail)
	return nil
}

// GetReasonDetail ReasonDetail Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetReasonDetail() string {
	return r._reasonDetail
}

// SetDealerOrderId is DealerOrderId Setter
// 经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// GetDealerOrderId DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

// SetReason is Reason Setter
// 关闭原因：&lt;br/&gt;1：长时间无法联系到分销商，取消交易。&lt;br/&gt;2：分销商错误提交申请，取消交易。&lt;br/&gt;3：缺货，近段时间都无法发货。&lt;br/&gt;4：分销商恶意提交申请单。&lt;br/&gt;5：其他原因。
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetReason(_reason int64) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetReason() int64 {
	return r._reason
}

var poolTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoDealerRequisitionorderCloseRequest()
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderCloseRequest 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest
func GetTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest() *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest {
	return poolTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest.Get().(*TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest 将 TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest(v *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderCloseAPIRequest.Put(v)
}
