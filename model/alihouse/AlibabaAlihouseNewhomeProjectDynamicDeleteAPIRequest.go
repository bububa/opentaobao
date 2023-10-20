package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest 楼盘快讯删除 API请求
// alibaba.alihouse.newhome.project.dynamic.delete
//
// 楼盘快讯删除
type AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest struct {
	model.Params
	// 外部动态ID
	_outerDynamicId string
	// 外部门店ID
	_outerStoreId string
}

// NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest 初始化AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest() *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest {
	return &AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) Reset() {
	r._outerDynamicId = ""
	r._outerStoreId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.dynamic.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterDynamicId is OuterDynamicId Setter
// 外部动态ID
func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) SetOuterDynamicId(_outerDynamicId string) error {
	r._outerDynamicId = _outerDynamicId
	r.Set("outer_dynamic_id", _outerDynamicId)
	return nil
}

// GetOuterDynamicId OuterDynamicId Getter
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetOuterDynamicId() string {
	return r._outerDynamicId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

var poolAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectDynamicDeleteRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectDynamicDeleteRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest
func GetAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest() *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest 将 AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest(v *AlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectDynamicDeleteAPIRequest.Put(v)
}
