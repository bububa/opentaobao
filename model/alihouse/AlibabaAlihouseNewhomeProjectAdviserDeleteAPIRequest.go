package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectadviserdeleteAPIRequest 删除楼盘顾问 API请求
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
type AlibabaalihousenewhomeprojectadviserdeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerConsultantId string
	// 外部门店ID
	_outerStoreId string
	// 版本号，请求时间戳
	_version int64
}

// NewAlibabaalihousenewhomeprojectadviserdeleteRequest 初始化AlibabaalihousenewhomeprojectadviserdeleteAPIRequest对象
func NewAlibabaalihousenewhomeprojectadviserdeleteRequest() *AlibabaalihousenewhomeprojectadviserdeleteAPIRequest {
	return &AlibabaalihousenewhomeprojectadviserdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterConsultantId is OuterConsultantId Setter
// 外部顾问ID
func (r *AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) SetOuterConsultantId(_outerConsultantId string) error {
	r._outerConsultantId = _outerConsultantId
	r.Set("outer_consultant_id", _outerConsultantId)
	return nil
}

// GetOuterConsultantId OuterConsultantId Getter
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetOuterConsultantId() string {
	return r._outerConsultantId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetVersion is Version Setter
// 版本号，请求时间戳
func (r *AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaalihousenewhomeprojectadviserdeleteAPIRequest) GetVersion() int64 {
	return r._version
}
