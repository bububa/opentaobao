package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasSendStatusAPIRequest 淘宝客-推广者-红包领取状态查询 API请求
// taobao.tbk.dg.vegas.send.status
//
// 淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkDgVegasSendStatusAPIRequest struct {
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

// NewTaobaoTbkDgVegasSendStatusRequest 初始化TaobaoTbkDgVegasSendStatusAPIRequest对象
func NewTaobaoTbkDgVegasSendStatusRequest() *TaobaoTbkDgVegasSendStatusAPIRequest {
	return &TaobaoTbkDgVegasSendStatusAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) Reset() {
	r._relationId = ""
	r._specialId = ""
	r._deviceValue = ""
	r._deviceType = ""
	r._thorBizCode = ""
	r._pid = ""
	r._activityCategory = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationId is RelationId Setter
// 渠道管理id
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetDeviceValue is DeviceValue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetThorBizCode is ThorBizCode Setter
// 已废弃，请勿传入
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetThorBizCode(_thorBizCode string) error {
	r._thorBizCode = _thorBizCode
	r.Set("thor_biz_code", _thorBizCode)
	return nil
}

// GetThorBizCode ThorBizCode Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetThorBizCode() string {
	return r._thorBizCode
}

// SetPid is Pid Setter
// 媒体pid
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetPid() string {
	return r._pid
}

// SetActivityCategory is ActivityCategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetActivityCategory(_activityCategory int64) error {
	r._activityCategory = _activityCategory
	r.Set("activity_category", _activityCategory)
	return nil
}

// GetActivityCategory ActivityCategory Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetActivityCategory() int64 {
	return r._activityCategory
}

var poolTaobaoTbkDgVegasSendStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgVegasSendStatusRequest()
	},
}

// GetTaobaoTbkDgVegasSendStatusRequest 从 sync.Pool 获取 TaobaoTbkDgVegasSendStatusAPIRequest
func GetTaobaoTbkDgVegasSendStatusAPIRequest() *TaobaoTbkDgVegasSendStatusAPIRequest {
	return poolTaobaoTbkDgVegasSendStatusAPIRequest.Get().(*TaobaoTbkDgVegasSendStatusAPIRequest)
}

// ReleaseTaobaoTbkDgVegasSendStatusAPIRequest 将 TaobaoTbkDgVegasSendStatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgVegasSendStatusAPIRequest(v *TaobaoTbkDgVegasSendStatusAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgVegasSendStatusAPIRequest.Put(v)
}
