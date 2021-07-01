package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest
供应商/分销商通过采购申请/经销采购单申请 API请求
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核 */
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest struct {
	model.Params
	// 采购申请/经销采购单编号
	_dealerOrderId int64
}

// NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest 初始化TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest() *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DealerOrderId Setter
// 采购申请/经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// Get DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}
