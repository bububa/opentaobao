package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest 提交楼盘电话 API请求
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest struct {
	model.Params
	// 外部门店ID
	_outerStoreId string
	// 楼盘电话
	_projectPhoneDto *ProjectPhoneDto
}

// NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest 初始化AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest() *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) Reset() {
	r._outerStoreId = ""
	r._projectPhoneDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.phone.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetProjectPhoneDto is ProjectPhoneDto Setter
// 楼盘电话
func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) SetProjectPhoneDto(_projectPhoneDto *ProjectPhoneDto) error {
	r._projectPhoneDto = _projectPhoneDto
	r.Set("project_phone_dto", _projectPhoneDto)
	return nil
}

// GetProjectPhoneDto ProjectPhoneDto Getter
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) GetProjectPhoneDto() *ProjectPhoneDto {
	return r._projectPhoneDto
}

var poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectPhoneSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIRequest.Put(v)
}
