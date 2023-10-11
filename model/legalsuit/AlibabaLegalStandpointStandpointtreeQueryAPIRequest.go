package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointtreeQueryAPIRequest 查询口径树目录 API请求
// alibaba.legal.standpoint.standpointtree.query
//
// 查询口径树目录
type AlibabaLegalStandpointStandpointtreeQueryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 场景id
	_sceneId int64
}

// NewAlibabaLegalStandpointStandpointtreeQueryRequest 初始化AlibabaLegalStandpointStandpointtreeQueryAPIRequest对象
func NewAlibabaLegalStandpointStandpointtreeQueryRequest() *AlibabaLegalStandpointStandpointtreeQueryAPIRequest {
	return &AlibabaLegalStandpointStandpointtreeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointStandpointtreeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpointtree.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointStandpointtreeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointStandpointtreeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointStandpointtreeQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointStandpointtreeQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetSceneId is SceneId Setter
// 场景id
func (r *AlibabaLegalStandpointStandpointtreeQueryAPIRequest) SetSceneId(_sceneId int64) error {
	r._sceneId = _sceneId
	r.Set("scene_id", _sceneId)
	return nil
}

// GetSceneId SceneId Getter
func (r AlibabaLegalStandpointStandpointtreeQueryAPIRequest) GetSceneId() int64 {
	return r._sceneId
}
