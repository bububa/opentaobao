package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivitySaveAPIRequest 新增或者更新销售活动 API请求
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
type AlibabaAlihouseNewhomeActivitySaveAPIRequest struct {
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

// NewAlibabaAlihouseNewhomeActivitySaveRequest 初始化AlibabaAlihouseNewhomeActivitySaveAPIRequest对象
func NewAlibabaAlihouseNewhomeActivitySaveRequest() *AlibabaAlihouseNewhomeActivitySaveAPIRequest {
	return &AlibabaAlihouseNewhomeActivitySaveAPIRequest{
		Params: model.NewParams(18),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) Reset() {
	r._outerSalesActivityId = ""
	r._outerProjectId = ""
	r._outerStoreId = ""
	r._activityName = ""
	r._activityTypeDesc = ""
	r._prov = ""
	r._provId = ""
	r._city = ""
	r._cityId = ""
	r._area = ""
	r._areaId = ""
	r._detailAddress = ""
	r._reminderContent = ""
	r._activityType = 0
	r._enablePreDepositGold = 0
	r._enableCustomer = 0
	r._isDeleted = 0
	r._isTest = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterSalesActivityId is OuterSalesActivityId Setter
// 外部销售活动id
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetOuterSalesActivityId(_outerSalesActivityId string) error {
	r._outerSalesActivityId = _outerSalesActivityId
	r.Set("outer_sales_activity_id", _outerSalesActivityId)
	return nil
}

// GetOuterSalesActivityId OuterSalesActivityId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetOuterSalesActivityId() string {
	return r._outerSalesActivityId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部项目店id
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetActivityTypeDesc is ActivityTypeDesc Setter
// 活动描述
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetActivityTypeDesc(_activityTypeDesc string) error {
	r._activityTypeDesc = _activityTypeDesc
	r.Set("activity_type_desc", _activityTypeDesc)
	return nil
}

// GetActivityTypeDesc ActivityTypeDesc Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetActivityTypeDesc() string {
	return r._activityTypeDesc
}

// SetProv is Prov Setter
// 省名称
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetProv(_prov string) error {
	r._prov = _prov
	r.Set("prov", _prov)
	return nil
}

// GetProv Prov Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetProv() string {
	return r._prov
}

// SetProvId is ProvId Setter
// 省编码
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetProvId(_provId string) error {
	r._provId = _provId
	r.Set("prov_id", _provId)
	return nil
}

// GetProvId ProvId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetProvId() string {
	return r._provId
}

// SetCity is City Setter
// 市名称
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetCity() string {
	return r._city
}

// SetCityId is CityId Setter
// 市编码
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetCityId(_cityId string) error {
	r._cityId = _cityId
	r.Set("city_id", _cityId)
	return nil
}

// GetCityId CityId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetCityId() string {
	return r._cityId
}

// SetArea is Area Setter
// 区名称
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetArea() string {
	return r._area
}

// SetAreaId is AreaId Setter
// 区编码
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetAreaId(_areaId string) error {
	r._areaId = _areaId
	r.Set("area_id", _areaId)
	return nil
}

// GetAreaId AreaId Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetAreaId() string {
	return r._areaId
}

// SetDetailAddress is DetailAddress Setter
// 详细地址
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetDetailAddress(_detailAddress string) error {
	r._detailAddress = _detailAddress
	r.Set("detail_address", _detailAddress)
	return nil
}

// GetDetailAddress DetailAddress Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetDetailAddress() string {
	return r._detailAddress
}

// SetReminderContent is ReminderContent Setter
// 商家须知
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetReminderContent(_reminderContent string) error {
	r._reminderContent = _reminderContent
	r.Set("reminder_content", _reminderContent)
	return nil
}

// GetReminderContent ReminderContent Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetReminderContent() string {
	return r._reminderContent
}

// SetActivityType is ActivityType Setter
// 活动类型(1-开盘活动 2-顺销活动)
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetActivityType(_activityType int64) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetActivityType() int64 {
	return r._activityType
}

// SetEnablePreDepositGold is EnablePreDepositGold Setter
// 是否启用预存金
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetEnablePreDepositGold(_enablePreDepositGold int64) error {
	r._enablePreDepositGold = _enablePreDepositGold
	r.Set("enable_pre_deposit_gold", _enablePreDepositGold)
	return nil
}

// GetEnablePreDepositGold EnablePreDepositGold Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetEnablePreDepositGold() int64 {
	return r._enablePreDepositGold
}

// SetEnableCustomer is EnableCustomer Setter
// 是否启用定向客户选房功能
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetEnableCustomer(_enableCustomer int64) error {
	r._enableCustomer = _enableCustomer
	r.Set("enable_customer", _enableCustomer)
	return nil
}

// GetEnableCustomer EnableCustomer Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetEnableCustomer() int64 {
	return r._enableCustomer
}

// SetIsDeleted is IsDeleted Setter
// 是否删除
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetIsDeleted(_isDeleted int64) error {
	r._isDeleted = _isDeleted
	r.Set("is_deleted", _isDeleted)
	return nil
}

// GetIsDeleted IsDeleted Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetIsDeleted() int64 {
	return r._isDeleted
}

// SetIsTest is IsTest Setter
// 是否测试
func (r *AlibabaAlihouseNewhomeActivitySaveAPIRequest) SetIsTest(_isTest int64) error {
	r._isTest = _isTest
	r.Set("is_test", _isTest)
	return nil
}

// GetIsTest IsTest Getter
func (r AlibabaAlihouseNewhomeActivitySaveAPIRequest) GetIsTest() int64 {
	return r._isTest
}

var poolAlibabaAlihouseNewhomeActivitySaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeActivitySaveRequest()
	},
}

// GetAlibabaAlihouseNewhomeActivitySaveRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivitySaveAPIRequest
func GetAlibabaAlihouseNewhomeActivitySaveAPIRequest() *AlibabaAlihouseNewhomeActivitySaveAPIRequest {
	return poolAlibabaAlihouseNewhomeActivitySaveAPIRequest.Get().(*AlibabaAlihouseNewhomeActivitySaveAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeActivitySaveAPIRequest 将 AlibabaAlihouseNewhomeActivitySaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivitySaveAPIRequest(v *AlibabaAlihouseNewhomeActivitySaveAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivitySaveAPIRequest.Put(v)
}
