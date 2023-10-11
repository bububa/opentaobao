package legalsuit

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
