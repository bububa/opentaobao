package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionorderagreeAPIRequest 供应商/分销商通过采购申请/经销采购单申请 API请求
// taobao.fenxiao.dealer.requisitionorder.agree
//
// 供应商或分销商通过采购申请/经销采购单审核
type TaobaofenxiaodealerrequisitionorderagreeAPIRequest struct {
	model.Params
	// 采购申请/经销采购单编号
	_dealerOrderId int64
}

// NewTaobaofenxiaodealerrequisitionorderagreeRequest 初始化TaobaofenxiaodealerrequisitionorderagreeAPIRequest对象
func NewTaobaofenxiaodealerrequisitionorderagreeRequest() *TaobaofenxiaodealerrequisitionorderagreeAPIRequest {
	return &TaobaofenxiaodealerrequisitionorderagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodealerrequisitionorderagreeAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodealerrequisitionorderagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodealerrequisitionorderagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDealerOrderId is DealerOrderId Setter
// 采购申请/经销采购单编号
func (r *TaobaofenxiaodealerrequisitionorderagreeAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// GetDealerOrderId DealerOrderId Getter
func (r TaobaofenxiaodealerrequisitionorderagreeAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}
