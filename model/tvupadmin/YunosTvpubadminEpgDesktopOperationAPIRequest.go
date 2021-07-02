package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminEpgDesktopOperationAPIRequest 影视桌面运营通用接口 API请求
// yunos.tvpubadmin.epg.desktop.operation
//
// 影视桌面运营通用接口
type YunosTvpubadminEpgDesktopOperationAPIRequest struct {
	model.Params
	// 操作对象实体
	_entityType string
	// 操作类型
	_actionType string
	// 具体入参
	_parameter string
}

// NewYunosTvpubadminEpgDesktopOperationRequest 初始化YunosTvpubadminEpgDesktopOperationAPIRequest对象
func NewYunosTvpubadminEpgDesktopOperationRequest() *YunosTvpubadminEpgDesktopOperationAPIRequest {
	return &YunosTvpubadminEpgDesktopOperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminEpgDesktopOperationAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.epg.desktop.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminEpgDesktopOperationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEntityType is EntityType Setter
// 操作对象实体
func (r *YunosTvpubadminEpgDesktopOperationAPIRequest) SetEntityType(_entityType string) error {
	r._entityType = _entityType
	r.Set("entity_type", _entityType)
	return nil
}

// GetEntityType EntityType Getter
func (r YunosTvpubadminEpgDesktopOperationAPIRequest) GetEntityType() string {
	return r._entityType
}

// SetActionType is ActionType Setter
// 操作类型
func (r *YunosTvpubadminEpgDesktopOperationAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r YunosTvpubadminEpgDesktopOperationAPIRequest) GetActionType() string {
	return r._actionType
}

// SetParameter is Parameter Setter
// 具体入参
func (r *YunosTvpubadminEpgDesktopOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunosTvpubadminEpgDesktopOperationAPIRequest) GetParameter() string {
	return r._parameter
}
