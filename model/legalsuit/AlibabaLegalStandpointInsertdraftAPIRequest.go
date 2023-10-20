package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointInsertdraftAPIRequest 插入草稿 API请求
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
type AlibabaLegalStandpointInsertdraftAPIRequest struct {
	model.Params
	// busid
	_busid string
	// busName
	_busName string
	// 系统标识
	_inputSystemCode string
	// 口径对象
	_standpointDraftDto *StandpointDraftDto
}

// NewAlibabaLegalStandpointInsertdraftRequest 初始化AlibabaLegalStandpointInsertdraftAPIRequest对象
func NewAlibabaLegalStandpointInsertdraftRequest() *AlibabaLegalStandpointInsertdraftAPIRequest {
	return &AlibabaLegalStandpointInsertdraftAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointInsertdraftAPIRequest) Reset() {
	r._busid = ""
	r._busName = ""
	r._inputSystemCode = ""
	r._standpointDraftDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.insertdraft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusid is Busid Setter
// busid
func (r *AlibabaLegalStandpointInsertdraftAPIRequest) SetBusid(_busid string) error {
	r._busid = _busid
	r.Set("busid", _busid)
	return nil
}

// GetBusid Busid Getter
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetBusid() string {
	return r._busid
}

// SetBusName is BusName Setter
// busName
func (r *AlibabaLegalStandpointInsertdraftAPIRequest) SetBusName(_busName string) error {
	r._busName = _busName
	r.Set("bus_name", _busName)
	return nil
}

// GetBusName BusName Getter
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetBusName() string {
	return r._busName
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointInsertdraftAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointDraftDto is StandpointDraftDto Setter
// 口径对象
func (r *AlibabaLegalStandpointInsertdraftAPIRequest) SetStandpointDraftDto(_standpointDraftDto *StandpointDraftDto) error {
	r._standpointDraftDto = _standpointDraftDto
	r.Set("standpoint_draft_dto", _standpointDraftDto)
	return nil
}

// GetStandpointDraftDto StandpointDraftDto Getter
func (r AlibabaLegalStandpointInsertdraftAPIRequest) GetStandpointDraftDto() *StandpointDraftDto {
	return r._standpointDraftDto
}

var poolAlibabaLegalStandpointInsertdraftAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointInsertdraftRequest()
	},
}

// GetAlibabaLegalStandpointInsertdraftRequest 从 sync.Pool 获取 AlibabaLegalStandpointInsertdraftAPIRequest
func GetAlibabaLegalStandpointInsertdraftAPIRequest() *AlibabaLegalStandpointInsertdraftAPIRequest {
	return poolAlibabaLegalStandpointInsertdraftAPIRequest.Get().(*AlibabaLegalStandpointInsertdraftAPIRequest)
}

// ReleaseAlibabaLegalStandpointInsertdraftAPIRequest 将 AlibabaLegalStandpointInsertdraftAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointInsertdraftAPIRequest(v *AlibabaLegalStandpointInsertdraftAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointInsertdraftAPIRequest.Put(v)
}
