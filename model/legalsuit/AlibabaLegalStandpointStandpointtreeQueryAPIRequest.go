package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointstandpointtreequeryAPIRequest 查询口径树目录 API请求
// alibaba.legal.standpoint.standpointtree.query
//
// 查询口径树目录
type AlibabalegalstandpointstandpointtreequeryAPIRequest struct {
	model.Params
	// 系统标识
	_inputSystemCode string
	// 场景id
	_sceneId int64
}

// NewAlibabalegalstandpointstandpointtreequeryRequest 初始化AlibabalegalstandpointstandpointtreequeryAPIRequest对象
func NewAlibabalegalstandpointstandpointtreequeryRequest() *AlibabalegalstandpointstandpointtreequeryAPIRequest {
	return &AlibabalegalstandpointstandpointtreequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointstandpointtreequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.standpointtree.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointstandpointtreequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointstandpointtreequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointstandpointtreequeryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointstandpointtreequeryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetSceneId is SceneId Setter
// 场景id
func (r *AlibabalegalstandpointstandpointtreequeryAPIRequest) SetSceneId(_sceneId int64) error {
	r._sceneId = _sceneId
	r.Set("scene_id", _sceneId)
	return nil
}

// GetSceneId SceneId Getter
func (r AlibabalegalstandpointstandpointtreequeryAPIRequest) GetSceneId() int64 {
	return r._sceneId
}
