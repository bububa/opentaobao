package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionordercloseAPIRequest 供应商/分销商关闭采购申请/经销采购单 API请求
// taobao.fenxiao.dealer.requisitionorder.close
//
// 供应商或分销商关闭采购申请/经销采购单
type TaobaofenxiaodealerrequisitionordercloseAPIRequest struct {
	model.Params
	// 关闭详细原因，字数5-200字
	_reasonDetail string
	// 经销采购单编号
	_dealerOrderId int64
	// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
	_reason int64
}

// NewTaobaofenxiaodealerrequisitionordercloseRequest 初始化TaobaofenxiaodealerrequisitionordercloseAPIRequest对象
func NewTaobaofenxiaodealerrequisitionordercloseRequest() *TaobaofenxiaodealerrequisitionordercloseAPIRequest {
	return &TaobaofenxiaodealerrequisitionordercloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReasonDetail is ReasonDetail Setter
// 关闭详细原因，字数5-200字
func (r *TaobaofenxiaodealerrequisitionordercloseAPIRequest) SetReasonDetail(_reasonDetail string) error {
	r._reasonDetail = _reasonDetail
	r.Set("reason_detail", _reasonDetail)
	return nil
}

// GetReasonDetail ReasonDetail Getter
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetReasonDetail() string {
	return r._reasonDetail
}

// SetDealerOrderId is DealerOrderId Setter
// 经销采购单编号
func (r *TaobaofenxiaodealerrequisitionordercloseAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// GetDealerOrderId DealerOrderId Getter
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

// SetReason is Reason Setter
// 关闭原因：&lt;br/&gt;1：长时间无法联系到分销商，取消交易。&lt;br/&gt;2：分销商错误提交申请，取消交易。&lt;br/&gt;3：缺货，近段时间都无法发货。&lt;br/&gt;4：分销商恶意提交申请单。&lt;br/&gt;5：其他原因。
func (r *TaobaofenxiaodealerrequisitionordercloseAPIRequest) SetReason(_reason int64) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaofenxiaodealerrequisitionordercloseAPIRequest) GetReason() int64 {
	return r._reason
}
