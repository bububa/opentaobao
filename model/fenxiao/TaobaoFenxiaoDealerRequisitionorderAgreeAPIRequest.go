package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest 供应商/分销商通过采购申请/经销采购单申请 API请求
// taobao.fenxiao.dealer.requisitionorder.agree
//
// 供应商或分销商通过采购申请/经销采购单审核
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest struct {
	model.Params
	// 采购申请/经销采购单编号
	_dealerOrderId int64
}

// NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest 初始化TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest() *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) Reset() {
	r._dealerOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDealerOrderId is DealerOrderId Setter
// 采购申请/经销采购单编号
func (r *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// GetDealerOrderId DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

var poolTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoDealerRequisitionorderAgreeRequest()
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderAgreeRequest 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest
func GetTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest() *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest {
	return poolTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest.Get().(*TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest 将 TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest(v *TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest.Put(v)
}
