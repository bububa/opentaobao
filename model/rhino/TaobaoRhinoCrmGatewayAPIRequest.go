package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoCrmGatewayAPIRequest crm实体变更回调接口 API请求
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
type TaobaoRhinoCrmGatewayAPIRequest struct {
	model.Params
	// 探迹定义，formcode
	_formCode string
	// 操作类型，create，update，delete
	_operateType string
	// 事件发生的时间戳
	_eventTimestamp int64
	// crm实体
	_crmEntity *CrmEntity
}

// NewTaobaoRhinoCrmGatewayRequest 初始化TaobaoRhinoCrmGatewayAPIRequest对象
func NewTaobaoRhinoCrmGatewayRequest() *TaobaoRhinoCrmGatewayAPIRequest {
	return &TaobaoRhinoCrmGatewayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRhinoCrmGatewayAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.crm.gateway"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRhinoCrmGatewayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRhinoCrmGatewayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFormCode is FormCode Setter
// 探迹定义，formcode
func (r *TaobaoRhinoCrmGatewayAPIRequest) SetFormCode(_formCode string) error {
	r._formCode = _formCode
	r.Set("form_code", _formCode)
	return nil
}

// GetFormCode FormCode Getter
func (r TaobaoRhinoCrmGatewayAPIRequest) GetFormCode() string {
	return r._formCode
}

// SetOperateType is OperateType Setter
// 操作类型，create，update，delete
func (r *TaobaoRhinoCrmGatewayAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoRhinoCrmGatewayAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetEventTimestamp is EventTimestamp Setter
// 事件发生的时间戳
func (r *TaobaoRhinoCrmGatewayAPIRequest) SetEventTimestamp(_eventTimestamp int64) error {
	r._eventTimestamp = _eventTimestamp
	r.Set("event_timestamp", _eventTimestamp)
	return nil
}

// GetEventTimestamp EventTimestamp Getter
func (r TaobaoRhinoCrmGatewayAPIRequest) GetEventTimestamp() int64 {
	return r._eventTimestamp
}

// SetCrmEntity is CrmEntity Setter
// crm实体
func (r *TaobaoRhinoCrmGatewayAPIRequest) SetCrmEntity(_crmEntity *CrmEntity) error {
	r._crmEntity = _crmEntity
	r.Set("crm_entity", _crmEntity)
	return nil
}

// GetCrmEntity CrmEntity Getter
func (r TaobaoRhinoCrmGatewayAPIRequest) GetCrmEntity() *CrmEntity {
	return r._crmEntity
}
