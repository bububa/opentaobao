package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecommunitylineAPIRequest 小区上下架 API请求
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
type AlibabaalihousenewhomecommunitylineAPIRequest struct {
	model.Params
	// 外部小区ID
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaalihousenewhomecommunitylineRequest 初始化AlibabaalihousenewhomecommunitylineAPIRequest对象
func NewAlibabaalihousenewhomecommunitylineRequest() *AlibabaalihousenewhomecommunitylineAPIRequest {
	return &AlibabaalihousenewhomecommunitylineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomecommunitylineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomecommunitylineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomecommunitylineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部小区ID
func (r *AlibabaalihousenewhomecommunitylineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomecommunitylineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaalihousenewhomecommunitylineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihousenewhomecommunitylineAPIRequest) GetType() *model.File {
	return r._type
}
