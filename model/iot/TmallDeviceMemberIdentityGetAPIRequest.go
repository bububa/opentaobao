package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallDeviceMemberIdentityGetAPIRequest 智能硬件会员判断 API请求
// tmall.device.member.identity.get
//
// 用来识别该用户是否是商家会员·
type TmallDeviceMemberIdentityGetAPIRequest struct {
	model.Params
	// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
	_extraInfo string
	// 混淆昵称，
	_mixNick string
	// 明文nick，可不填，直接填混淆昵称
	_nick string
}

// NewTmallDeviceMemberIdentityGetRequest 初始化TmallDeviceMemberIdentityGetAPIRequest对象
func NewTmallDeviceMemberIdentityGetRequest() *TmallDeviceMemberIdentityGetAPIRequest {
	return &TmallDeviceMemberIdentityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallDeviceMemberIdentityGetAPIRequest) GetApiMethodName() string {
	return "tmall.device.member.identity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallDeviceMemberIdentityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtraInfo Setter
// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
func (r *TmallDeviceMemberIdentityGetAPIRequest) SetExtraInfo(_extraInfo string) error {
	r._extraInfo = _extraInfo
	r.Set("extra_info", _extraInfo)
	return nil
}

// Get ExtraInfo Getter
func (r TmallDeviceMemberIdentityGetAPIRequest) GetExtraInfo() string {
	return r._extraInfo
}

// Set is MixNick Setter
// 混淆昵称，
func (r *TmallDeviceMemberIdentityGetAPIRequest) SetMixNick(_mixNick string) error {
	r._mixNick = _mixNick
	r.Set("mix_nick", _mixNick)
	return nil
}

// Get MixNick Getter
func (r TmallDeviceMemberIdentityGetAPIRequest) GetMixNick() string {
	return r._mixNick
}

// Set is Nick Setter
// 明文nick，可不填，直接填混淆昵称
func (r *TmallDeviceMemberIdentityGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TmallDeviceMemberIdentityGetAPIRequest) GetNick() string {
	return r._nick
}
