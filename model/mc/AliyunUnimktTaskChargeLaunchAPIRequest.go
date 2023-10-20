package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyununimkttaskchargelaunchAPIRequest 云码权益查询 API请求
// aliyun.unimkt.task.charge.launch
//
// 云码线上流量投放链路，用于判断用户是否有匹配的投放计划
type AliyununimkttaskchargelaunchAPIRequest struct {
	model.Params
	// 服务商附加url参数
	_extra string
	// urlID
	_urlId string
	// 支付宝openID
	_alipayOpenId string
	// 渠道ID
	_channelId string
	// 淘宝ID
	_userId string
}

// NewAliyununimkttaskchargelaunchRequest 初始化AliyununimkttaskchargelaunchAPIRequest对象
func NewAliyununimkttaskchargelaunchRequest() *AliyununimkttaskchargelaunchAPIRequest {
	return &AliyununimkttaskchargelaunchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyununimkttaskchargelaunchAPIRequest) GetApiMethodName() string {
	return "aliyun.unimkt.task.charge.launch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyununimkttaskchargelaunchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyununimkttaskchargelaunchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtra is Extra Setter
// 服务商附加url参数
func (r *AliyununimkttaskchargelaunchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AliyununimkttaskchargelaunchAPIRequest) GetExtra() string {
	return r._extra
}

// SetUrlId is UrlId Setter
// urlID
func (r *AliyununimkttaskchargelaunchAPIRequest) SetUrlId(_urlId string) error {
	r._urlId = _urlId
	r.Set("url_id", _urlId)
	return nil
}

// GetUrlId UrlId Getter
func (r AliyununimkttaskchargelaunchAPIRequest) GetUrlId() string {
	return r._urlId
}

// SetAlipayOpenId is AlipayOpenId Setter
// 支付宝openID
func (r *AliyununimkttaskchargelaunchAPIRequest) SetAlipayOpenId(_alipayOpenId string) error {
	r._alipayOpenId = _alipayOpenId
	r.Set("alipay_open_id", _alipayOpenId)
	return nil
}

// GetAlipayOpenId AlipayOpenId Getter
func (r AliyununimkttaskchargelaunchAPIRequest) GetAlipayOpenId() string {
	return r._alipayOpenId
}

// SetChannelId is ChannelId Setter
// 渠道ID
func (r *AliyununimkttaskchargelaunchAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AliyununimkttaskchargelaunchAPIRequest) GetChannelId() string {
	return r._channelId
}

// SetUserId is UserId Setter
// 淘宝ID
func (r *AliyununimkttaskchargelaunchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AliyununimkttaskchargelaunchAPIRequest) GetUserId() string {
	return r._userId
}
