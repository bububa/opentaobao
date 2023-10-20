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
	_relationid string
	// 会员运营id
	_specialid string
	// 加密后的值，需用MD5加密，32位小写
	_devicevalue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
	_devicetype string
	// 已废弃，请勿传入
	_thorbizcode string
	// 媒体pid
	_pid string
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activitycategory int64
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

// SetRelationid is Relationid Setter
// 渠道管理id
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetSpecialid is Specialid Setter
// 会员运营id
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetDevicevalue is Devicevalue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetDevicevalue(_devicevalue string) error {
	r._devicevalue = _devicevalue
	r.Set("device_value", _devicevalue)
	return nil
}

// GetDevicevalue Devicevalue Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetDevicevalue() string {
	return r._devicevalue
}

// SetDevicetype is Devicetype Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetDevicetype(_devicetype string) error {
	r._devicetype = _devicetype
	r.Set("device_type", _devicetype)
	return nil
}

// GetDevicetype Devicetype Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetDevicetype() string {
	return r._devicetype
}

// SetThorbizcode is Thorbizcode Setter
// 已废弃，请勿传入
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetThorbizcode(_thorbizcode string) error {
	r._thorbizcode = _thorbizcode
	r.Set("thor_biz_code", _thorbizcode)
	return nil
}

// GetThorbizcode Thorbizcode Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetThorbizcode() string {
	return r._thorbizcode
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

// SetActivitycategory is Activitycategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkdgvegassendstatusAPIRequest) SetActivitycategory(_activitycategory int64) error {
	r._activitycategory = _activitycategory
	r.Set("activity_category", _activitycategory)
	return nil
}

// GetActivitycategory Activitycategory Getter
func (r TaobaotbkdgvegassendstatusAPIRequest) GetActivitycategory() int64 {
	return r._activitycategory
}
