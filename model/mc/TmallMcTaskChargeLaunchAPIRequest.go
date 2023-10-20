package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmctaskchargelaunchAPIRequest 云码充电宝投放链路 API请求
// tmall.mc.task.charge.launch
//
// 云码充电宝投放链路，用于判断用户是否有匹配的投放计划
type TmallmctaskchargelaunchAPIRequest struct {
	model.Params
	// 外部设备编码
	_outerCode string
	// 渠道ID
	_channelId string
	// 支付宝openID
	_alipayOpenId string
	// urlID
	_urlId string
	// 服务商附加url参数
	_extra string
}

// NewTmallmctaskchargelaunchRequest 初始化TmallmctaskchargelaunchAPIRequest对象
func NewTmallmctaskchargelaunchRequest() *TmallmctaskchargelaunchAPIRequest {
	return &TmallmctaskchargelaunchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmctaskchargelaunchAPIRequest) GetApiMethodName() string {
	return "tmall.mc.task.charge.launch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmctaskchargelaunchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmctaskchargelaunchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 外部设备编码
func (r *TmallmctaskchargelaunchAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TmallmctaskchargelaunchAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetChannelId is ChannelId Setter
// 渠道ID
func (r *TmallmctaskchargelaunchAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TmallmctaskchargelaunchAPIRequest) GetChannelId() string {
	return r._channelId
}

// SetAlipayOpenId is AlipayOpenId Setter
// 支付宝openID
func (r *TmallmctaskchargelaunchAPIRequest) SetAlipayOpenId(_alipayOpenId string) error {
	r._alipayOpenId = _alipayOpenId
	r.Set("alipay_open_id", _alipayOpenId)
	return nil
}

// GetAlipayOpenId AlipayOpenId Getter
func (r TmallmctaskchargelaunchAPIRequest) GetAlipayOpenId() string {
	return r._alipayOpenId
}

// SetUrlId is UrlId Setter
// urlID
func (r *TmallmctaskchargelaunchAPIRequest) SetUrlId(_urlId string) error {
	r._urlId = _urlId
	r.Set("url_id", _urlId)
	return nil
}

// GetUrlId UrlId Getter
func (r TmallmctaskchargelaunchAPIRequest) GetUrlId() string {
	return r._urlId
}

// SetExtra is Extra Setter
// 服务商附加url参数
func (r *TmallmctaskchargelaunchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TmallmctaskchargelaunchAPIRequest) GetExtra() string {
	return r._extra
}
