package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScVegasSendStatusAPIRequest 淘宝客-服务商-红包领取状态查询 API请求
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
type TaobaoTbkScVegasSendStatusAPIRequest struct {
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

// NewTaobaoTbkScVegasSendStatusRequest 初始化TaobaoTbkScVegasSendStatusAPIRequest对象
func NewTaobaoTbkScVegasSendStatusRequest() *TaobaoTbkScVegasSendStatusAPIRequest {
	return &TaobaoTbkScVegasSendStatusAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScVegasSendStatusAPIRequest) Reset() {
	r._specialId = ""
	r._relationId = ""
	r._deviceValue = ""
	r._deviceType = ""
	r._activityCategory = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetRelationId is RelationId Setter
// 渠道管理id
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetDeviceValue is DeviceValue Setter
// 加密后的值，需用MD5加密，32位小写
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetActivityCategory is ActivityCategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetActivityCategory(_activityCategory int64) error {
	r._activityCategory = _activityCategory
	r.Set("activity_category", _activityCategory)
	return nil
}

// GetActivityCategory ActivityCategory Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetActivityCategory() int64 {
	return r._activityCategory
}

var poolTaobaoTbkScVegasSendStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScVegasSendStatusRequest()
	},
}

// GetTaobaoTbkScVegasSendStatusRequest 从 sync.Pool 获取 TaobaoTbkScVegasSendStatusAPIRequest
func GetTaobaoTbkScVegasSendStatusAPIRequest() *TaobaoTbkScVegasSendStatusAPIRequest {
	return poolTaobaoTbkScVegasSendStatusAPIRequest.Get().(*TaobaoTbkScVegasSendStatusAPIRequest)
}

// ReleaseTaobaoTbkScVegasSendStatusAPIRequest 将 TaobaoTbkScVegasSendStatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScVegasSendStatusAPIRequest(v *TaobaoTbkScVegasSendStatusAPIRequest) {
	v.Reset()
	poolTaobaoTbkScVegasSendStatusAPIRequest.Put(v)
}
