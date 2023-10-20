package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeagreementpreshowAPIRequest 预览地址获取接口 API请求
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
type AlibabaalihousenewhomeagreementpreshowAPIRequest struct {
	model.Params
	// 协议ID
	_id int64
	// 1-协议 2-签章
	_type int64
}

// NewAlibabaalihousenewhomeagreementpreshowRequest 初始化AlibabaalihousenewhomeagreementpreshowAPIRequest对象
func NewAlibabaalihousenewhomeagreementpreshowRequest() *AlibabaalihousenewhomeagreementpreshowAPIRequest {
	return &AlibabaalihousenewhomeagreementpreshowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeagreementpreshowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.agreement.preshow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeagreementpreshowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeagreementpreshowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 协议ID
func (r *AlibabaalihousenewhomeagreementpreshowAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaalihousenewhomeagreementpreshowAPIRequest) GetId() int64 {
	return r._id
}

// SetType is Type Setter
// 1-协议 2-签章
func (r *AlibabaalihousenewhomeagreementpreshowAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihousenewhomeagreementpreshowAPIRequest) GetType() int64 {
	return r._type
}
