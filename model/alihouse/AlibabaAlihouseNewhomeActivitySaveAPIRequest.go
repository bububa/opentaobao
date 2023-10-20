package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitysaveAPIRequest 新增或者更新销售活动 API请求
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
type AlibabaalihousenewhomeactivitysaveAPIRequest struct {
	model.Params
	// 外部销售活动id
	_outerSalesActivityId string
	// 外部楼盘id
	_outerProjectId string
	// 外部项目店id
	_outerStoreId string
	// 活动名称
	_activityName string
	// 活动描述
	_activityTypeDesc string
	// 省名称
	_prov string
	// 省编码
	_provId string
	// 市名称
	_city string
	// 市编码
	_cityId string
	// 区名称
	_area string
	// 区编码
	_areaId string
	// 详细地址
	_detailAddress string
	// 商家须知
	_reminderContent string
	// 活动类型(1-开盘活动 2-顺销活动)
	_activityType int64
	// 是否启用预存金
	_enablePreDepositGold int64
	// 是否启用定向客户选房功能
	_enableCustomer int64
	// 是否删除
	_isDeleted int64
	// 是否测试
	_isTest int64
}

// NewAlibabaalihousenewhomeactivitysaveRequest 初始化AlibabaalihousenewhomeactivitysaveAPIRequest对象
func NewAlibabaalihousenewhomeactivitysaveRequest() *AlibabaalihousenewhomeactivitysaveAPIRequest {
	return &AlibabaalihousenewhomeactivitysaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterSalesActivityId is OuterSalesActivityId Setter
// 外部销售活动id
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetOuterSalesActivityId(_outerSalesActivityId string) error {
	r._outerSalesActivityId = _outerSalesActivityId
	r.Set("outer_sales_activity_id", _outerSalesActivityId)
	return nil
}

// GetOuterSalesActivityId OuterSalesActivityId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetOuterSalesActivityId() string {
	return r._outerSalesActivityId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部项目店id
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetActivityTypeDesc is ActivityTypeDesc Setter
// 活动描述
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetActivityTypeDesc(_activityTypeDesc string) error {
	r._activityTypeDesc = _activityTypeDesc
	r.Set("activity_type_desc", _activityTypeDesc)
	return nil
}

// GetActivityTypeDesc ActivityTypeDesc Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetActivityTypeDesc() string {
	return r._activityTypeDesc
}

// SetProv is Prov Setter
// 省名称
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetProv(_prov string) error {
	r._prov = _prov
	r.Set("prov", _prov)
	return nil
}

// GetProv Prov Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetProv() string {
	return r._prov
}

// SetProvId is ProvId Setter
// 省编码
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetProvId(_provId string) error {
	r._provId = _provId
	r.Set("prov_id", _provId)
	return nil
}

// GetProvId ProvId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetProvId() string {
	return r._provId
}

// SetCity is City Setter
// 市名称
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetCity() string {
	return r._city
}

// SetCityId is CityId Setter
// 市编码
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetCityId(_cityId string) error {
	r._cityId = _cityId
	r.Set("city_id", _cityId)
	return nil
}

// GetCityId CityId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetCityId() string {
	return r._cityId
}

// SetArea is Area Setter
// 区名称
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetArea() string {
	return r._area
}

// SetAreaId is AreaId Setter
// 区编码
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetAreaId(_areaId string) error {
	r._areaId = _areaId
	r.Set("area_id", _areaId)
	return nil
}

// GetAreaId AreaId Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetAreaId() string {
	return r._areaId
}

// SetDetailAddress is DetailAddress Setter
// 详细地址
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetDetailAddress(_detailAddress string) error {
	r._detailAddress = _detailAddress
	r.Set("detail_address", _detailAddress)
	return nil
}

// GetDetailAddress DetailAddress Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetDetailAddress() string {
	return r._detailAddress
}

// SetReminderContent is ReminderContent Setter
// 商家须知
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetReminderContent(_reminderContent string) error {
	r._reminderContent = _reminderContent
	r.Set("reminder_content", _reminderContent)
	return nil
}

// GetReminderContent ReminderContent Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetReminderContent() string {
	return r._reminderContent
}

// SetActivityType is ActivityType Setter
// 活动类型(1-开盘活动 2-顺销活动)
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetActivityType(_activityType int64) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetActivityType() int64 {
	return r._activityType
}

// SetEnablePreDepositGold is EnablePreDepositGold Setter
// 是否启用预存金
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetEnablePreDepositGold(_enablePreDepositGold int64) error {
	r._enablePreDepositGold = _enablePreDepositGold
	r.Set("enable_pre_deposit_gold", _enablePreDepositGold)
	return nil
}

// GetEnablePreDepositGold EnablePreDepositGold Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetEnablePreDepositGold() int64 {
	return r._enablePreDepositGold
}

// SetEnableCustomer is EnableCustomer Setter
// 是否启用定向客户选房功能
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetEnableCustomer(_enableCustomer int64) error {
	r._enableCustomer = _enableCustomer
	r.Set("enable_customer", _enableCustomer)
	return nil
}

// GetEnableCustomer EnableCustomer Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetEnableCustomer() int64 {
	return r._enableCustomer
}

// SetIsDeleted is IsDeleted Setter
// 是否删除
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetIsDeleted(_isDeleted int64) error {
	r._isDeleted = _isDeleted
	r.Set("is_deleted", _isDeleted)
	return nil
}

// GetIsDeleted IsDeleted Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetIsDeleted() int64 {
	return r._isDeleted
}

// SetIsTest is IsTest Setter
// 是否测试
func (r *AlibabaalihousenewhomeactivitysaveAPIRequest) SetIsTest(_isTest int64) error {
	r._isTest = _isTest
	r.Set("is_test", _isTest)
	return nil
}

// GetIsTest IsTest Getter
func (r AlibabaalihousenewhomeactivitysaveAPIRequest) GetIsTest() int64 {
	return r._isTest
}
