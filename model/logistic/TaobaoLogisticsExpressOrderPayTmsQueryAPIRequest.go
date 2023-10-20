package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressorderpaytmsqueryAPIRequest 上门取退运费支付状态查询接口 API请求
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
type TaobaologisticsexpressorderpaytmsqueryAPIRequest struct {
	model.Params
	// 查询入参
	_tms2MscPayQueryRequest *Tms2mscPayQueryRequest
}

// NewTaobaologisticsexpressorderpaytmsqueryRequest 初始化TaobaologisticsexpressorderpaytmsqueryAPIRequest对象
func NewTaobaologisticsexpressorderpaytmsqueryRequest() *TaobaologisticsexpressorderpaytmsqueryAPIRequest {
	return &TaobaologisticsexpressorderpaytmsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressorderpaytmsqueryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.pay.tms.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressorderpaytmsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressorderpaytmsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTms2MscPayQueryRequest is Tms2MscPayQueryRequest Setter
// 查询入参
func (r *TaobaologisticsexpressorderpaytmsqueryAPIRequest) SetTms2MscPayQueryRequest(_tms2MscPayQueryRequest *Tms2mscPayQueryRequest) error {
	r._tms2MscPayQueryRequest = _tms2MscPayQueryRequest
	r.Set("tms2_msc_pay_query_request", _tms2MscPayQueryRequest)
	return nil
}

// GetTms2MscPayQueryRequest Tms2MscPayQueryRequest Getter
func (r TaobaologisticsexpressorderpaytmsqueryAPIRequest) GetTms2MscPayQueryRequest() *Tms2mscPayQueryRequest {
	return r._tms2MscPayQueryRequest
}
