package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest 顾问入职离职 API请求
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
type AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerConsultantId string
	// 外部门店id
	_outerShopId string
	// 1-入职 2-离职
	_status int64
	// 版本号，时间戳
	_version int64
}

// NewAlibabaAlihouseNewhomeAdviserSubmitAccountRequest 初始化AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest对象
func NewAlibabaAlihouseNewhomeAdviserSubmitAccountRequest() *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest {
	return &AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) Reset() {
	r._outerConsultantId = ""
	r._outerShopId = ""
	r._status = 0
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.adviser.submit.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterConsultantId is OuterConsultantId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) SetOuterConsultantId(_outerConsultantId string) error {
	r._outerConsultantId = _outerConsultantId
	r.Set("outer_consultant_id", _outerConsultantId)
	return nil
}

// GetOuterConsultantId OuterConsultantId Getter
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetOuterConsultantId() string {
	return r._outerConsultantId
}

// SetOuterShopId is OuterShopId Setter
// 外部门店id
func (r *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetStatus is Status Setter
// 1-入职 2-离职
func (r *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetStatus() int64 {
	return r._status
}

// SetVersion is Version Setter
// 版本号，时间戳
func (r *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeAdviserSubmitAccountRequest()
	},
}

// GetAlibabaAlihouseNewhomeAdviserSubmitAccountRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest
func GetAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest() *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest {
	return poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest.Get().(*AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest 将 AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest(v *AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest.Put(v)
}
