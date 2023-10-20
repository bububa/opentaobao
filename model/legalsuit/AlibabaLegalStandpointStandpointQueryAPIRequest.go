package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointQueryAPIRequest 查询具体口径 API请求
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
type AlibabaLegalStandpointStandpointQueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 口径id
	_standpointId int64
}

// NewAlibabaLegalStandpointStandpointQueryRequest 初始化AlibabaLegalStandpointStandpointQueryAPIRequest对象
func NewAlibabaLegalStandpointStandpointQueryRequest() *AlibabaLegalStandpointStandpointQueryAPIRequest {
	return &AlibabaLegalStandpointStandpointQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointStandpointQueryAPIRequest) Reset() {
	r._inputSystemCode = ""
	r._standpointId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointStandpointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointStandpointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointStandpointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointStandpointQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointStandpointQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetStandpointId is StandpointId Setter
// 口径id
func (r *AlibabaLegalStandpointStandpointQueryAPIRequest) SetStandpointId(_standpointId int64) error {
	r._standpointId = _standpointId
	r.Set("standpoint_id", _standpointId)
	return nil
}

// GetStandpointId StandpointId Getter
func (r AlibabaLegalStandpointStandpointQueryAPIRequest) GetStandpointId() int64 {
	return r._standpointId
}

var poolAlibabaLegalStandpointStandpointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointStandpointQueryRequest()
	},
}

// GetAlibabaLegalStandpointStandpointQueryRequest 从 sync.Pool 获取 AlibabaLegalStandpointStandpointQueryAPIRequest
func GetAlibabaLegalStandpointStandpointQueryAPIRequest() *AlibabaLegalStandpointStandpointQueryAPIRequest {
	return poolAlibabaLegalStandpointStandpointQueryAPIRequest.Get().(*AlibabaLegalStandpointStandpointQueryAPIRequest)
}

// ReleaseAlibabaLegalStandpointStandpointQueryAPIRequest 将 AlibabaLegalStandpointStandpointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointStandpointQueryAPIRequest(v *AlibabaLegalStandpointStandpointQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointStandpointQueryAPIRequest.Put(v)
}
