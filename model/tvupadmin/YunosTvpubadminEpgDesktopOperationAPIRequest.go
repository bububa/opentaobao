package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminepgdesktopoperationAPIRequest 影视桌面运营通用接口 API请求
// yunos.tvpubadmin.epg.desktop.operation
//
// 影视桌面运营通用接口
type YunostvpubadminepgdesktopoperationAPIRequest struct {
	model.Params
	// 操作对象实体
	_entityType string
	// 操作类型
	_actionType string
	// 具体入参
	_parameter string
}

// NewYunostvpubadminepgdesktopoperationRequest 初始化YunostvpubadminepgdesktopoperationAPIRequest对象
func NewYunostvpubadminepgdesktopoperationRequest() *YunostvpubadminepgdesktopoperationAPIRequest {
	return &YunostvpubadminepgdesktopoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.epg.desktop.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntityType is EntityType Setter
// 操作对象实体
func (r *YunostvpubadminepgdesktopoperationAPIRequest) SetEntityType(_entityType string) error {
	r._entityType = _entityType
	r.Set("entity_type", _entityType)
	return nil
}

// GetEntityType EntityType Getter
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetEntityType() string {
	return r._entityType
}

// SetActionType is ActionType Setter
// 操作类型
func (r *YunostvpubadminepgdesktopoperationAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetActionType() string {
	return r._actionType
}

// SetParameter is Parameter Setter
// 具体入参
func (r *YunostvpubadminepgdesktopoperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunostvpubadminepgdesktopoperationAPIRequest) GetParameter() string {
	return r._parameter
}
