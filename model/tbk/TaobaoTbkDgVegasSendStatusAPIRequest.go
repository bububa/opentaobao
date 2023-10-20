package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegassendstatusAPIRequest 淘宝客-推广者-红包领取状态查询 API请求
// taobao.tbk.dg.vegas.send.status
//
// 淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkdgvegassendstatusAPIRequest struct {
	model.Params
	// 渠道管理id
	_relationId string
	// 会员运营id
	_specialId string
	// 加密后的值，需用MD5加密，32位小写
	_deviceValue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
	_deviceType string
	// 已废弃，请勿传入
	_thorBizCode string
	// 媒体pid
	_pid string
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activityCategory int64
}

// NewTaobaotbkdgvegassendstatusRequest 初始化TaobaotbkdgvegassendstatusAPIRequest对象
func NewTaobaotbkdgvegassendstatusRequest() *TaobaotbkdgvegassendstatusAPIRequest {
	return &TaobaotbkdgvegassendstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgvegassendstatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgvegassendstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgvegassendstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationId is RelationId Setter
// 渠道管理id
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetDeviceValue is DeviceValue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetThorBizCode is ThorBizCode Setter
// 已废弃，请勿传入
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetThorBizCode(_thorBizCode string) error {
	r._thorBizCode = _thorBizCode
	r.Set("thor_biz_code", _thorBizCode)
	return nil
}

// GetThorBizCode ThorBizCode Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetThorBizCode() string {
	return r._thorBizCode
}

// SetPid is Pid Setter
// 媒体pid
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetPid() string {
	return r._pid
}

// SetActivityCategory is ActivityCategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetActivityCategory(_activityCategory int64) error {
	r._activityCategory = _activityCategory
	r.Set("activity_category", _activityCategory)
	return nil
}

// GetActivityCategory ActivityCategory Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetActivityCategory() int64 {
	return r._activityCategory
}
