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
	_specialid string
	// 渠道管理id
	_relationid string
	// 加密后的值，需用MD5加密，32位小写
	_devicevalue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
	_devicetype string
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activitycategory int64
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

// SetSpecialid is Specialid Setter
// 会员运营id
func (r *TaobaotbkscvegassendstatusAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetRelationid is Relationid Setter
// 渠道管理id
func (r *TaobaotbkscvegassendstatusAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetDevicevalue is Devicevalue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaotbkscvegassendstatusAPIRequest) SetDevicevalue(_devicevalue string) error {
	r._devicevalue = _devicevalue
	r.Set("device_value", _devicevalue)
	return nil
}

// GetDevicevalue Devicevalue Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetDevicevalue() string {
	return r._devicevalue
}

// SetDevicetype is Devicetype Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaotbkscvegassendstatusAPIRequest) SetDevicetype(_devicetype string) error {
	r._devicetype = _devicetype
	r.Set("device_type", _devicetype)
	return nil
}

// GetDevicetype Devicetype Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetDevicetype() string {
	return r._devicetype
}

// SetActivitycategory is Activitycategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkscvegassendstatusAPIRequest) SetActivitycategory(_activitycategory int64) error {
	r._activitycategory = _activitycategory
	r.Set("activity_category", _activitycategory)
	return nil
}

// GetActivitycategory Activitycategory Getter
func (r TaobaotbkscvegassendstatusAPIRequest) GetActivitycategory() int64 {
	return r._activitycategory
}
