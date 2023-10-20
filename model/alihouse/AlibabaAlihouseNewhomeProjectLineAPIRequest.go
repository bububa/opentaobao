package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectLineAPIRequest 楼盘上下架 API请求
// alibaba.alihouse.newhome.project.line
//
// 上下架楼盘
type AlibabaAlihouseNewhomeProjectLineAPIRequest struct {
	model.Params
	// 外部门店ID
	_outerStoreId string
	// 外部楼盘ID
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaAlihouseNewhomeProjectLineRequest 初始化AlibabaAlihouseNewhomeProjectLineAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectLineRequest() *AlibabaAlihouseNewhomeProjectLineAPIRequest {
	return &AlibabaAlihouseNewhomeProjectLineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaAlihouseNewhomeProjectLineAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetOuterId is OuterId Setter
// 外部楼盘ID
func (r *AlibabaAlihouseNewhomeProjectLineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaAlihouseNewhomeProjectLineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetType() *model.File {
	return r._type
}
