package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointDeleteAPIRequest 删除关联口径 API请求
// alibaba.legal.standpoint.delete
//
// 删除关联口径
type AlibabaLegalStandpointDeleteAPIRequest struct {
	model.Params
	// 工号
	_userId string
	// 系统标识
	_inputSystemCode string
	// 关联id
	_refrenceId int64
}

// NewAlibabaLegalStandpointDeleteRequest 初始化AlibabaLegalStandpointDeleteAPIRequest对象
func NewAlibabaLegalStandpointDeleteRequest() *AlibabaLegalStandpointDeleteAPIRequest {
	return &AlibabaLegalStandpointDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointDeleteAPIRequest) Reset() {
	r._userId = ""
	r._inputSystemCode = ""
	r._refrenceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 工号
func (r *AlibabaLegalStandpointDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLegalStandpointDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointDeleteAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointDeleteAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetRefrenceId is RefrenceId Setter
// 关联id
func (r *AlibabaLegalStandpointDeleteAPIRequest) SetRefrenceId(_refrenceId int64) error {
	r._refrenceId = _refrenceId
	r.Set("refrence_id", _refrenceId)
	return nil
}

// GetRefrenceId RefrenceId Getter
func (r AlibabaLegalStandpointDeleteAPIRequest) GetRefrenceId() int64 {
	return r._refrenceId
}

var poolAlibabaLegalStandpointDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointDeleteRequest()
	},
}

// GetAlibabaLegalStandpointDeleteRequest 从 sync.Pool 获取 AlibabaLegalStandpointDeleteAPIRequest
func GetAlibabaLegalStandpointDeleteAPIRequest() *AlibabaLegalStandpointDeleteAPIRequest {
	return poolAlibabaLegalStandpointDeleteAPIRequest.Get().(*AlibabaLegalStandpointDeleteAPIRequest)
}

// ReleaseAlibabaLegalStandpointDeleteAPIRequest 将 AlibabaLegalStandpointDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointDeleteAPIRequest(v *AlibabaLegalStandpointDeleteAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointDeleteAPIRequest.Put(v)
}
