package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardverifycoderesendAPIRequest 重发核销码 API请求
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
type TmallservicecenterworkcardverifycoderesendAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 门店/网点id
	_serviceStoreId int64
}

// NewTmallservicecenterworkcardverifycoderesendRequest 初始化TmallservicecenterworkcardverifycoderesendAPIRequest对象
func NewTmallservicecenterworkcardverifycoderesendRequest() *TmallservicecenterworkcardverifycoderesendAPIRequest {
	return &TmallservicecenterworkcardverifycoderesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardverifycoderesendAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.verifycode.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardverifycoderesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardverifycoderesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardverifycoderesendAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardverifycoderesendAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetServiceStoreId is ServiceStoreId Setter
// 门店/网点id
func (r *TmallservicecenterworkcardverifycoderesendAPIRequest) SetServiceStoreId(_serviceStoreId int64) error {
	r._serviceStoreId = _serviceStoreId
	r.Set("service_store_id", _serviceStoreId)
	return nil
}

// GetServiceStoreId ServiceStoreId Getter
func (r TmallservicecenterworkcardverifycoderesendAPIRequest) GetServiceStoreId() int64 {
	return r._serviceStoreId
}
