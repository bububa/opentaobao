package mc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMcTaskChargeLaunchAPIRequest 云码充电宝投放链路 API请求
// tmall.mc.task.charge.launch
//
// 云码充电宝投放链路，用于判断用户是否有匹配的投放计划
type TmallMcTaskChargeLaunchAPIRequest struct {
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

// NewTmallMcTaskChargeLaunchRequest 初始化TmallMcTaskChargeLaunchAPIRequest对象
func NewTmallMcTaskChargeLaunchRequest() *TmallMcTaskChargeLaunchAPIRequest {
	return &TmallMcTaskChargeLaunchAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMcTaskChargeLaunchAPIRequest) Reset() {
	r._outerCode = ""
	r._channelId = ""
	r._alipayOpenId = ""
	r._urlId = ""
	r._extra = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMcTaskChargeLaunchAPIRequest) GetApiMethodName() string {
	return "tmall.mc.task.charge.launch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMcTaskChargeLaunchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMcTaskChargeLaunchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// 外部设备编码
func (r *TmallMcTaskChargeLaunchAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TmallMcTaskChargeLaunchAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetChannelId is ChannelId Setter
// 渠道ID
func (r *TmallMcTaskChargeLaunchAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TmallMcTaskChargeLaunchAPIRequest) GetChannelId() string {
	return r._channelId
}

// SetAlipayOpenId is AlipayOpenId Setter
// 支付宝openID
func (r *TmallMcTaskChargeLaunchAPIRequest) SetAlipayOpenId(_alipayOpenId string) error {
	r._alipayOpenId = _alipayOpenId
	r.Set("alipay_open_id", _alipayOpenId)
	return nil
}

// GetAlipayOpenId AlipayOpenId Getter
func (r TmallMcTaskChargeLaunchAPIRequest) GetAlipayOpenId() string {
	return r._alipayOpenId
}

// SetUrlId is UrlId Setter
// urlID
func (r *TmallMcTaskChargeLaunchAPIRequest) SetUrlId(_urlId string) error {
	r._urlId = _urlId
	r.Set("url_id", _urlId)
	return nil
}

// GetUrlId UrlId Getter
func (r TmallMcTaskChargeLaunchAPIRequest) GetUrlId() string {
	return r._urlId
}

// SetExtra is Extra Setter
// 服务商附加url参数
func (r *TmallMcTaskChargeLaunchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TmallMcTaskChargeLaunchAPIRequest) GetExtra() string {
	return r._extra
}

var poolTmallMcTaskChargeLaunchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMcTaskChargeLaunchRequest()
	},
}

// GetTmallMcTaskChargeLaunchRequest 从 sync.Pool 获取 TmallMcTaskChargeLaunchAPIRequest
func GetTmallMcTaskChargeLaunchAPIRequest() *TmallMcTaskChargeLaunchAPIRequest {
	return poolTmallMcTaskChargeLaunchAPIRequest.Get().(*TmallMcTaskChargeLaunchAPIRequest)
}

// ReleaseTmallMcTaskChargeLaunchAPIRequest 将 TmallMcTaskChargeLaunchAPIRequest 放入 sync.Pool
func ReleaseTmallMcTaskChargeLaunchAPIRequest(v *TmallMcTaskChargeLaunchAPIRequest) {
	v.Reset()
	poolTmallMcTaskChargeLaunchAPIRequest.Put(v)
}
