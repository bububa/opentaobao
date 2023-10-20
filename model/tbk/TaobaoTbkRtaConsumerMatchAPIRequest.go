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
	_offerlist []OfferList
	// 消费者对应的会员ID（会员ID或设备信息同时填时，优先使用会员ID）
	_specialid string
	// 设备信息，加密后的值(IMEI,IDFA,OAID,MOBILE需要加密)，需用MD5加密，32位小写
	_devicevalue string
	// 设备信息，入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)：IMEI, 或者IDFA, 或者OAID, 或者MOBILE, 或者ALIPAY_ID
	_devicetype string
	// 策略ID，与活动列表二选一传入
	_strategyidlist string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneid int64
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

// SetOfferlist is Offerlist Setter
// 活动列表
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetOfferlist(_offerlist []OfferList) error {
	r._offerlist = _offerlist
	r.Set("offer_list", _offerlist)
	return nil
}

// GetOfferlist Offerlist Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetOfferlist() []OfferList {
	return r._offerlist
}

// SetSpecialid is Specialid Setter
// 消费者对应的会员ID（会员ID或设备信息同时填时，优先使用会员ID）
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetDevicevalue is Devicevalue Setter
// 设备信息，加密后的值(IMEI,IDFA,OAID,MOBILE需要加密)，需用MD5加密，32位小写
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetDevicevalue(_devicevalue string) error {
	r._devicevalue = _devicevalue
	r.Set("device_value", _devicevalue)
	return nil
}

// GetDevicevalue Devicevalue Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetDevicevalue() string {
	return r._devicevalue
}

// SetDevicetype is Devicetype Setter
// 设备信息，入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)：IMEI, 或者IDFA, 或者OAID, 或者MOBILE, 或者ALIPAY_ID
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetDevicetype(_devicetype string) error {
	r._devicetype = _devicetype
	r.Set("device_type", _devicetype)
	return nil
}

// GetDevicetype Devicetype Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetDevicetype() string {
	return r._devicetype
}

// SetStrategyidlist is Strategyidlist Setter
// 策略ID，与活动列表二选一传入
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetStrategyidlist(_strategyidlist string) error {
	r._strategyidlist = _strategyidlist
	r.Set("strategy_id_list", _strategyidlist)
	return nil
}

// GetStrategyidlist Strategyidlist Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetStrategyidlist() string {
	return r._strategyidlist
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkrtaconsumermatchAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkrtaconsumermatchAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}
