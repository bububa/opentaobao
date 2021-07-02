package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasSendStatusAPIRequest 淘宝客-推广者-超级红包领取状态查询 API请求
// taobao.tbk.dg.vegas.send.status
//
// 淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkDgVegasSendStatusAPIRequest struct {
	model.Params
	// 渠道管理id
	_relationId string
	// 会员运营id
	_specialId string
	// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
	_deviceValue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
	_deviceType string
	// thor平台业务码， 1：coupon 超红
	_thorBizCode string
	// 媒体pid
	_pid string
}

// NewTaobaoTbkDgVegasSendStatusRequest 初始化TaobaoTbkDgVegasSendStatusAPIRequest对象
func NewTaobaoTbkDgVegasSendStatusRequest() *TaobaoTbkDgVegasSendStatusAPIRequest {
	return &TaobaoTbkDgVegasSendStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RelationId Setter
// 渠道管理id
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// Get RelationId Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// Set is SpecialId Setter
// 会员运营id
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// Get SpecialId Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// Set is DeviceValue Setter
// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// Get DeviceValue Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// Set is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// Get DeviceType Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// Set is ThorBizCode Setter
// thor平台业务码， 1：coupon 超红
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetThorBizCode(_thorBizCode string) error {
	r._thorBizCode = _thorBizCode
	r.Set("thor_biz_code", _thorBizCode)
	return nil
}

// Get ThorBizCode Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetThorBizCode() string {
	return r._thorBizCode
}

// Set is Pid Setter
// 媒体pid
func (r *TaobaoTbkDgVegasSendStatusAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// Get Pid Getter
func (r TaobaoTbkDgVegasSendStatusAPIRequest) GetPid() string {
	return r._pid
}
