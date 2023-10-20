package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsorderQueryAPIRequest 物流单信息查询 API请求
// tmall.servicecenter.workcard.logisticsorder.query
//
// 物流订单信息查询API
type TmallServicecenterWorkcardLogisticsorderQueryAPIRequest struct {
	model.Params
	// 物流单号
	_logisticsOrderId int64
	// 是否查询新物流，不填默认查原有物流
	_newLogistics bool
}

// NewTmallServicecenterWorkcardLogisticsorderQueryRequest 初始化TmallServicecenterWorkcardLogisticsorderQueryAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsorderQueryRequest() *TmallServicecenterWorkcardLogisticsorderQueryAPIRequest {
	return &TmallServicecenterWorkcardLogisticsorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流单号
func (r *TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}

// SetNewLogistics is NewLogistics Setter
// 是否查询新物流，不填默认查原有物流
func (r *TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) SetNewLogistics(_newLogistics bool) error {
	r._newLogistics = _newLogistics
	r.Set("new_logistics", _newLogistics)
	return nil
}

// GetNewLogistics NewLogistics Getter
func (r TmallServicecenterWorkcardLogisticsorderQueryAPIRequest) GetNewLogistics() bool {
	return r._newLogistics
}
