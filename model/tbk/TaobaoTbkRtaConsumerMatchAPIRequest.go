package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkrtaconsumermatchAPIRequest 淘宝客-推广者-定向活动目标发布 API请求
// taobao.tbk.rta.consumer.match
//
// 淘客计划向用户推送某个定向活动时，调用该接口判断用户是否符合活动目标（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaotbkrtaconsumermatchAPIRequest struct {
	model.Params
	// 活动列表
	_offerList []OfferList
	// 消费者对应的会员ID（会员ID或设备信息同时填时，优先使用会员ID）
	_specialId string
	// 设备信息，加密后的值(IMEI,IDFA,OAID,MOBILE需要加密)，需用MD5加密，32位小写
	_deviceValue string
	// 设备信息，入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)：IMEI, 或者IDFA, 或者OAID, 或者MOBILE, 或者ALIPAY_ID
	_deviceType string
	// 策略ID，与活动列表二选一传入
	_strategyIdList string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
}

// NewTaobaotbkrtaconsumermatchRequest 初始化TaobaotbkrtaconsumermatchAPIRequest对象
func NewTaobaotbkrtaconsumermatchRequest() *TaobaotbkrtaconsumermatchAPIRequest {
	return &TaobaotbkrtaconsumermatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkrtaconsumermatchAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.rta.consumer.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkrtaconsumermatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkrtaconsumermatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfferList is OfferList Setter
// 活动列表
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetOfferList(_offerList []OfferList) error {
	r._offerList = _offerList
	r.Set("offer_list", _offerList)
	return nil
}

// GetOfferList OfferList Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetOfferList() []OfferList {
	return r._offerList
}

// SetSpecialId is SpecialId Setter
// 消费者对应的会员ID（会员ID或设备信息同时填时，优先使用会员ID）
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetDeviceValue is DeviceValue Setter
// 设备信息，加密后的值(IMEI,IDFA,OAID,MOBILE需要加密)，需用MD5加密，32位小写
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 设备信息，入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)：IMEI, 或者IDFA, 或者OAID, 或者MOBILE, 或者ALIPAY_ID
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetStrategyIdList is StrategyIdList Setter
// 策略ID，与活动列表二选一传入
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetStrategyIdList(_strategyIdList string) error {
	r._strategyIdList = _strategyIdList
	r.Set("strategy_id_list", _strategyIdList)
	return nil
}

// GetStrategyIdList StrategyIdList Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetStrategyIdList() string {
	return r._strategyIdList
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}
