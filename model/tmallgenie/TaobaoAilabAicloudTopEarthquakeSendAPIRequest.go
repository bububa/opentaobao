package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopearthquakesendAPIRequest 地震局发送地震消息 API请求
// taobao.ailab.aicloud.top.earthquake.send
//
// 地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
type TaobaoailabaicloudtopearthquakesendAPIRequest struct {
	model.Params
	// 扩展占位字段
	_ext string
	// 签名
	_signature string
	// 随机值
	_nonceStr string
	// 时间戳
	_timestampStr string
	// 地震信息
	_earthquakeInfo string
}

// NewTaobaoailabaicloudtopearthquakesendRequest 初始化TaobaoailabaicloudtopearthquakesendAPIRequest对象
func NewTaobaoailabaicloudtopearthquakesendRequest() *TaobaoailabaicloudtopearthquakesendAPIRequest {
	return &TaobaoailabaicloudtopearthquakesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.earthquake.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展占位字段
func (r *TaobaoailabaicloudtopearthquakesendAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetExt() string {
	return r._ext
}

// SetSignature is Signature Setter
// 签名
func (r *TaobaoailabaicloudtopearthquakesendAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetSignature() string {
	return r._signature
}

// SetNonceStr is NonceStr Setter
// 随机值
func (r *TaobaoailabaicloudtopearthquakesendAPIRequest) SetNonceStr(_nonceStr string) error {
	r._nonceStr = _nonceStr
	r.Set("nonce_str", _nonceStr)
	return nil
}

// GetNonceStr NonceStr Getter
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetNonceStr() string {
	return r._nonceStr
}

// SetTimestampStr is TimestampStr Setter
// 时间戳
func (r *TaobaoailabaicloudtopearthquakesendAPIRequest) SetTimestampStr(_timestampStr string) error {
	r._timestampStr = _timestampStr
	r.Set("timestamp_str", _timestampStr)
	return nil
}

// GetTimestampStr TimestampStr Getter
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetTimestampStr() string {
	return r._timestampStr
}

// SetEarthquakeInfo is EarthquakeInfo Setter
// 地震信息
func (r *TaobaoailabaicloudtopearthquakesendAPIRequest) SetEarthquakeInfo(_earthquakeInfo string) error {
	r._earthquakeInfo = _earthquakeInfo
	r.Set("earthquake_info", _earthquakeInfo)
	return nil
}

// GetEarthquakeInfo EarthquakeInfo Getter
func (r TaobaoailabaicloudtopearthquakesendAPIRequest) GetEarthquakeInfo() string {
	return r._earthquakeInfo
}
