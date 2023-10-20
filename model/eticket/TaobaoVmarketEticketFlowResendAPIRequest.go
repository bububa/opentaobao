package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketflowresendAPIRequest 业务重新触发发码短信 API请求
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
type TaobaovmarketeticketflowresendAPIRequest struct {
	model.Params
	// 业务单号
	_outerId string
	// 业务类型值，可联系淘宝业务运营取得具体值
	_bizType int64
}

// NewTaobaovmarketeticketflowresendRequest 初始化TaobaovmarketeticketflowresendAPIRequest对象
func NewTaobaovmarketeticketflowresendRequest() *TaobaovmarketeticketflowresendAPIRequest {
	return &TaobaovmarketeticketflowresendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketflowresendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.flow.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketflowresendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketflowresendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 业务单号
func (r *TaobaovmarketeticketflowresendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaovmarketeticketflowresendAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBizType is BizType Setter
// 业务类型值，可联系淘宝业务运营取得具体值
func (r *TaobaovmarketeticketflowresendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaovmarketeticketflowresendAPIRequest) GetBizType() int64 {
	return r._bizType
}
