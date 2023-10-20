package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointSceneQueryAPIRequest 查询场景 API请求
// alibaba.legal.standpoint.scene.query
//
// 查询场景
type AlibabaLegalStandpointSceneQueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
}

// NewAlibabaLegalStandpointSceneQueryRequest 初始化AlibabaLegalStandpointSceneQueryAPIRequest对象
func NewAlibabaLegalStandpointSceneQueryRequest() *AlibabaLegalStandpointSceneQueryAPIRequest {
	return &AlibabaLegalStandpointSceneQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointSceneQueryAPIRequest) Reset() {
	r._inputSystemCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointSceneQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.scene.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointSceneQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointSceneQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointSceneQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointSceneQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

var poolAlibabaLegalStandpointSceneQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointSceneQueryRequest()
	},
}

// GetAlibabaLegalStandpointSceneQueryRequest 从 sync.Pool 获取 AlibabaLegalStandpointSceneQueryAPIRequest
func GetAlibabaLegalStandpointSceneQueryAPIRequest() *AlibabaLegalStandpointSceneQueryAPIRequest {
	return poolAlibabaLegalStandpointSceneQueryAPIRequest.Get().(*AlibabaLegalStandpointSceneQueryAPIRequest)
}

// ReleaseAlibabaLegalStandpointSceneQueryAPIRequest 将 AlibabaLegalStandpointSceneQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointSceneQueryAPIRequest(v *AlibabaLegalStandpointSceneQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointSceneQueryAPIRequest.Put(v)
}
