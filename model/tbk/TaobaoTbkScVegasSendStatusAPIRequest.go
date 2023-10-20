package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscvegassendstatusAPIRequest 淘宝客-服务商-红包领取状态查询 API请求
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
type TaobaotbkscvegassendstatusAPIRequest struct {
	model.Params
	// 会员运营id
	_specialId string
	// 渠道管理id
	_relationId string
	// 加密后的值，需用MD5加密，32位小写
	_deviceValue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
	_deviceType string
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activityCategory int64
}

// NewTaobaotbkscvegassendstatusRequest 初始化TaobaotbkscvegassendstatusAPIRequest对象
func NewTaobaotbkscvegassendstatusRequest() *TaobaotbkscvegassendstatusAPIRequest {
	return &TaobaotbkscvegassendstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscvegassendstatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscvegassendstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscvegassendstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaotbkscvegassendstatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetRelationId is RelationId Setter
// 渠道管理id
func (r *TaobaotbkscvegassendstatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetDeviceValue is DeviceValue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaotbkscvegassendstatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaotbkscvegassendstatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetActivityCategory is ActivityCategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkscvegassendstatusAPIRequest) SetActivityCategory(_activityCategory int64) error {
	r._activityCategory = _activityCategory
	r.Set("activity_category", _activityCategory)
	return nil
}

// GetActivityCategory ActivityCategory Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetActivityCategory() int64 {
	return r._activityCategory
}
