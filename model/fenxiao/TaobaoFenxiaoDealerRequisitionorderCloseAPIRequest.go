package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest 供应商/分销商关闭采购申请/经销采购单 API请求
// taobao.fenxiao.dealer.requisitionorder.close
//
// 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest struct {
	model.Params
	// 经销采购单编号
	_dealerOrderId int64
	// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
	_reason int64
	// 关闭详细原因，字数5-200字
	_reasonDetail string
}

// NewTaobaoFenxiaoDealerRequisitionorderCloseRequest 初始化TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderCloseRequest() *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DealerOrderId Setter
// 经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// Get DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

// Set is Reason Setter
// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetReason(_reason int64) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetReason() int64 {
	return r._reason
}

// Set is ReasonDetail Setter
// 关闭详细原因，字数5-200字
func (r *TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) SetReasonDetail(_reasonDetail string) error {
	r._reasonDetail = _reasonDetail
	r.Set("reason_detail", _reasonDetail)
	return nil
}

// Get ReasonDetail Getter
func (r TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest) GetReasonDetail() string {
	return r._reasonDetail
}
