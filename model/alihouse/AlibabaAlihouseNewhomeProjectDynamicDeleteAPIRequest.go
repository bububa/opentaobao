package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest 楼盘快讯删除 API请求
// alibaba.alihouse.newhome.project.dynamic.delete
//
// 楼盘快讯删除
type AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest struct {
	model.Params
	// 外部动态ID
	_outerDynamicId string
	// 外部门店ID
	_outerStoreId string
}

// NewAlibabaalihousenewhomeprojectdynamicdeleteRequest 初始化AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest对象
func NewAlibabaalihousenewhomeprojectdynamicdeleteRequest() *AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest {
	return &AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.dynamic.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterDynamicId is OuterDynamicId Setter
// 外部动态ID
func (r *AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) SetOuterDynamicId(_outerDynamicId string) error {
	r._outerDynamicId = _outerDynamicId
	r.Set("outer_dynamic_id", _outerDynamicId)
	return nil
}

// GetOuterDynamicId OuterDynamicId Getter
func (r AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) GetOuterDynamicId() string {
	return r._outerDynamicId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}
