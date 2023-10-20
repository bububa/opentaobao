package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopEarthquakeSendAPIRequest 地震局发送地震消息 API请求
// taobao.ailab.aicloud.top.earthquake.send
//
// 地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
type TaobaoAilabAicloudTopEarthquakeSendAPIRequest struct {
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

// NewTaobaoAilabAicloudTopEarthquakeSendRequest 初始化TaobaoAilabAicloudTopEarthquakeSendAPIRequest对象
func NewTaobaoAilabAicloudTopEarthquakeSendRequest() *TaobaoAilabAicloudTopEarthquakeSendAPIRequest {
	return &TaobaoAilabAicloudTopEarthquakeSendAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) Reset() {
	r._ext = ""
	r._signature = ""
	r._nonceStr = ""
	r._timestampStr = ""
	r._earthquakeInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.earthquake.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展占位字段
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetExt() string {
	return r._ext
}

// SetSignature is Signature Setter
// 签名
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetSignature() string {
	return r._signature
}

// SetNonceStr is NonceStr Setter
// 随机值
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) SetNonceStr(_nonceStr string) error {
	r._nonceStr = _nonceStr
	r.Set("nonce_str", _nonceStr)
	return nil
}

// GetNonceStr NonceStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetNonceStr() string {
	return r._nonceStr
}

// SetTimestampStr is TimestampStr Setter
// 时间戳
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) SetTimestampStr(_timestampStr string) error {
	r._timestampStr = _timestampStr
	r.Set("timestamp_str", _timestampStr)
	return nil
}

// GetTimestampStr TimestampStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetTimestampStr() string {
	return r._timestampStr
}

// SetEarthquakeInfo is EarthquakeInfo Setter
// 地震信息
func (r *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) SetEarthquakeInfo(_earthquakeInfo string) error {
	r._earthquakeInfo = _earthquakeInfo
	r.Set("earthquake_info", _earthquakeInfo)
	return nil
}

// GetEarthquakeInfo EarthquakeInfo Getter
func (r TaobaoAilabAicloudTopEarthquakeSendAPIRequest) GetEarthquakeInfo() string {
	return r._earthquakeInfo
}

var poolTaobaoAilabAicloudTopEarthquakeSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopEarthquakeSendRequest()
	},
}

// GetTaobaoAilabAicloudTopEarthquakeSendRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopEarthquakeSendAPIRequest
func GetTaobaoAilabAicloudTopEarthquakeSendAPIRequest() *TaobaoAilabAicloudTopEarthquakeSendAPIRequest {
	return poolTaobaoAilabAicloudTopEarthquakeSendAPIRequest.Get().(*TaobaoAilabAicloudTopEarthquakeSendAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopEarthquakeSendAPIRequest 将 TaobaoAilabAicloudTopEarthquakeSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopEarthquakeSendAPIRequest(v *TaobaoAilabAicloudTopEarthquakeSendAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopEarthquakeSendAPIRequest.Put(v)
}
