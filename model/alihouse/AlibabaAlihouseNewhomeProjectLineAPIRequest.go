package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectlineAPIRequest 楼盘上下架 API请求
// alibaba.alihouse.newhome.project.line
//
// 上下架楼盘
type AlibabaalihousenewhomeprojectlineAPIRequest struct {
	model.Params
	// 外部门店ID
	_outerStoreId string
	// 外部楼盘ID
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaalihousenewhomeprojectlineRequest 初始化AlibabaalihousenewhomeprojectlineAPIRequest对象
func NewAlibabaalihousenewhomeprojectlineRequest() *AlibabaalihousenewhomeprojectlineAPIRequest {
	return &AlibabaalihousenewhomeprojectlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaalihousenewhomeprojectlineAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetOuterId is OuterId Setter
// 外部楼盘ID
func (r *AlibabaalihousenewhomeprojectlineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaalihousenewhomeprojectlineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihousenewhomeprojectlineAPIRequest) GetType() *model.File {
	return r._type
}
