package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinocrmgatewayAPIRequest crm实体变更回调接口 API请求
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
type TaobaorhinocrmgatewayAPIRequest struct {
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

// NewTaobaorhinocrmgatewayRequest 初始化TaobaorhinocrmgatewayAPIRequest对象
func NewTaobaorhinocrmgatewayRequest() *TaobaorhinocrmgatewayAPIRequest {
	return &TaobaorhinocrmgatewayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorhinocrmgatewayAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.crm.gateway"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorhinocrmgatewayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorhinocrmgatewayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFormCode is FormCode Setter
// 探迹定义，formcode
func (r *TaobaorhinocrmgatewayAPIRequest) SetFormCode(_formCode string) error {
	r._formCode = _formCode
	r.Set("form_code", _formCode)
	return nil
}

// GetFormCode FormCode Getter
func (r TaobaorhinocrmgatewayAPIRequest) GetFormCode() string {
	return r._formCode
}

// SetOperateType is OperateType Setter
// 操作类型，create，update，delete
func (r *TaobaorhinocrmgatewayAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaorhinocrmgatewayAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetEventTimestamp is EventTimestamp Setter
// 事件发生的时间戳
func (r *TaobaorhinocrmgatewayAPIRequest) SetEventTimestamp(_eventTimestamp int64) error {
	r._eventTimestamp = _eventTimestamp
	r.Set("event_timestamp", _eventTimestamp)
	return nil
}

// GetEventTimestamp EventTimestamp Getter
func (r TaobaorhinocrmgatewayAPIRequest) GetEventTimestamp() int64 {
	return r._eventTimestamp
}

// SetCrmEntity is CrmEntity Setter
// crm实体
func (r *TaobaorhinocrmgatewayAPIRequest) SetCrmEntity(_crmEntity *CrmEntity) error {
	r._crmEntity = _crmEntity
	r.Set("crm_entity", _crmEntity)
	return nil
}

// GetCrmEntity CrmEntity Getter
func (r TaobaorhinocrmgatewayAPIRequest) GetCrmEntity() *CrmEntity {
	return r._crmEntity
}
