package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunUnimktTaskChargeLaunchAPIRequest 云码权益查询 API请求
// aliyun.unimkt.task.charge.launch
//
// 云码线上流量投放链路，用于判断用户是否有匹配的投放计划
type AliyunUnimktTaskChargeLaunchAPIRequest struct {
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

// NewAliyunUnimktTaskChargeLaunchRequest 初始化AliyunUnimktTaskChargeLaunchAPIRequest对象
func NewAliyunUnimktTaskChargeLaunchRequest() *AliyunUnimktTaskChargeLaunchAPIRequest {
	return &AliyunUnimktTaskChargeLaunchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetApiMethodName() string {
	return "aliyun.unimkt.task.charge.launch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Extra Setter
// 服务商附加url参数
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// Get Extra Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetExtra() string {
	return r._extra
}

// Set is UrlId Setter
// urlID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetUrlId(_urlId string) error {
	r._urlId = _urlId
	r.Set("url_id", _urlId)
	return nil
}

// Get UrlId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetUrlId() string {
	return r._urlId
}

// Set is AlipayOpenId Setter
// 支付宝openID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetAlipayOpenId(_alipayOpenId string) error {
	r._alipayOpenId = _alipayOpenId
	r.Set("alipay_open_id", _alipayOpenId)
	return nil
}

// Get AlipayOpenId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetAlipayOpenId() string {
	return r._alipayOpenId
}

// Set is ChannelId Setter
// 渠道ID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// Get ChannelId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetChannelId() string {
	return r._channelId
}

// Set is UserId Setter
// 淘宝ID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetUserId() string {
	return r._userId
}
