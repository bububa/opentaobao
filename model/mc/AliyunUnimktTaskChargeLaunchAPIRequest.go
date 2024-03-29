package mc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) Reset() {
	r._extra = ""
	r._urlId = ""
	r._alipayOpenId = ""
	r._channelId = ""
	r._userId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetApiMethodName() string {
	return "aliyun.unimkt.task.charge.launch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtra is Extra Setter
// 服务商附加url参数
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetExtra() string {
	return r._extra
}

// SetUrlId is UrlId Setter
// urlID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetUrlId(_urlId string) error {
	r._urlId = _urlId
	r.Set("url_id", _urlId)
	return nil
}

// GetUrlId UrlId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetUrlId() string {
	return r._urlId
}

// SetAlipayOpenId is AlipayOpenId Setter
// 支付宝openID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetAlipayOpenId(_alipayOpenId string) error {
	r._alipayOpenId = _alipayOpenId
	r.Set("alipay_open_id", _alipayOpenId)
	return nil
}

// GetAlipayOpenId AlipayOpenId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetAlipayOpenId() string {
	return r._alipayOpenId
}

// SetChannelId is ChannelId Setter
// 渠道ID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetChannelId() string {
	return r._channelId
}

// SetUserId is UserId Setter
// 淘宝ID
func (r *AliyunUnimktTaskChargeLaunchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AliyunUnimktTaskChargeLaunchAPIRequest) GetUserId() string {
	return r._userId
}

var poolAliyunUnimktTaskChargeLaunchAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunUnimktTaskChargeLaunchRequest()
	},
}

// GetAliyunUnimktTaskChargeLaunchRequest 从 sync.Pool 获取 AliyunUnimktTaskChargeLaunchAPIRequest
func GetAliyunUnimktTaskChargeLaunchAPIRequest() *AliyunUnimktTaskChargeLaunchAPIRequest {
	return poolAliyunUnimktTaskChargeLaunchAPIRequest.Get().(*AliyunUnimktTaskChargeLaunchAPIRequest)
}

// ReleaseAliyunUnimktTaskChargeLaunchAPIRequest 将 AliyunUnimktTaskChargeLaunchAPIRequest 放入 sync.Pool
func ReleaseAliyunUnimktTaskChargeLaunchAPIRequest(v *AliyunUnimktTaskChargeLaunchAPIRequest) {
	v.Reset()
	poolAliyunUnimktTaskChargeLaunchAPIRequest.Put(v)
}
