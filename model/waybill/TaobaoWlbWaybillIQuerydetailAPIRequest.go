package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybilliquerydetailAPIRequest 查面单号状态v1.0 API请求
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
type TaobaowlbwaybilliquerydetailAPIRequest struct {
	model.Params
	// 面单查询请求
	_waybillDetailQueryRequest *WaybillDetailQueryRequest
}

// NewTaobaowlbwaybilliquerydetailRequest 初始化TaobaowlbwaybilliquerydetailAPIRequest对象
func NewTaobaowlbwaybilliquerydetailRequest() *TaobaowlbwaybilliquerydetailAPIRequest {
	return &TaobaowlbwaybilliquerydetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybilliquerydetailAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.querydetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybilliquerydetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybilliquerydetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillDetailQueryRequest is WaybillDetailQueryRequest Setter
// 面单查询请求
func (r *TaobaowlbwaybilliquerydetailAPIRequest) SetWaybillDetailQueryRequest(_waybillDetailQueryRequest *WaybillDetailQueryRequest) error {
	r._waybillDetailQueryRequest = _waybillDetailQueryRequest
	r.Set("waybill_detail_query_request", _waybillDetailQueryRequest)
	return nil
}

// GetWaybillDetailQueryRequest WaybillDetailQueryRequest Getter
func (r TaobaowlbwaybilliquerydetailAPIRequest) GetWaybillDetailQueryRequest() *WaybillDetailQueryRequest {
	return r._waybillDetailQueryRequest
}
