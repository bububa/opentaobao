package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardlogisticsorderqueryAPIRequest 物流单信息查询 API请求
// tmall.servicecenter.workcard.logisticsorder.query
//
// 物流订单信息查询API
type TmallservicecenterworkcardlogisticsorderqueryAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId int64
	// 是否查询新物流，不填默认查原有物流
	_newLogistics bool
}

// NewTmallservicecenterworkcardlogisticsorderqueryRequest 初始化TmallservicecenterworkcardlogisticsorderqueryAPIRequest对象
func NewTmallservicecenterworkcardlogisticsorderqueryRequest() *TmallservicecenterworkcardlogisticsorderqueryAPIRequest {
	return &TmallservicecenterworkcardlogisticsorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardlogisticsorderqueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardlogisticsorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardlogisticsorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流单号
func (r *TmallservicecenterworkcardlogisticsorderqueryAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r TmallservicecenterworkcardlogisticsorderqueryAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}

// SetNewLogistics is NewLogistics Setter
// 是否查询新物流，不填默认查原有物流
func (r *TmallservicecenterworkcardlogisticsorderqueryAPIRequest) SetNewLogistics(_newLogistics bool) error {
	r._newLogistics = _newLogistics
	r.Set("new_logistics", _newLogistics)
	return nil
}

// GetNewLogistics NewLogistics Getter
func (r TmallservicecenterworkcardlogisticsorderqueryAPIRequest) GetNewLogistics() bool {
	return r._newLogistics
}
