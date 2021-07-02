package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIQuerydetailAPIRequest 查面单号状态v1.0 API请求
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailAPIRequest struct {
	model.Params
	// 面单查询请求
	_waybillDetailQueryRequest *WaybillDetailQueryRequest
}

// NewTaobaoWlbWaybillIQuerydetailRequest 初始化TaobaoWlbWaybillIQuerydetailAPIRequest对象
func NewTaobaoWlbWaybillIQuerydetailRequest() *TaobaoWlbWaybillIQuerydetailAPIRequest {
	return &TaobaoWlbWaybillIQuerydetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WaybillDetailQueryRequest Setter
// 面单查询请求
func (r *TaobaoWlbWaybillIQuerydetailAPIRequest) SetWaybillDetailQueryRequest(_waybillDetailQueryRequest *WaybillDetailQueryRequest) error {
	r._waybillDetailQueryRequest = _waybillDetailQueryRequest
	r.Set("waybill_detail_query_request", _waybillDetailQueryRequest)
	return nil
}

// Get WaybillDetailQueryRequest Getter
func (r TaobaoWlbWaybillIQuerydetailAPIRequest) GetWaybillDetailQueryRequest() *WaybillDetailQueryRequest {
	return r._waybillDetailQueryRequest
}
