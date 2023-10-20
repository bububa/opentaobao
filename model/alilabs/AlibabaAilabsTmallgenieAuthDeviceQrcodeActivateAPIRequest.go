package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest 扫码激活设备 API请求
// alibaba.ailabs.tmallgenie.auth.device.qrcode.activate
//
// 三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。
type AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest struct {
	model.Params
	// OAUTH authcode码
	_code string
	// 产品终端ID
	_clientId string
	// 淘宝授权ID
	_taobaoUserOpenid string
	// 扩展数据
	_extInfo string
}

// NewAlibabaailabstmallgenieauthdeviceqrcodeactivateRequest 初始化AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest对象
func NewAlibabaailabstmallgenieauthdeviceqrcodeactivateRequest() *AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest {
	return &AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.qrcode.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// OAUTH authcode码
func (r *AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetCode() string {
	return r._code
}

// SetClientId is ClientId Setter
// 产品终端ID
func (r *AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetClientId() string {
	return r._clientId
}

// SetTaobaoUserOpenid is TaobaoUserOpenid Setter
// 淘宝授权ID
func (r *AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) SetTaobaoUserOpenid(_taobaoUserOpenid string) error {
	r._taobaoUserOpenid = _taobaoUserOpenid
	r.Set("taobao_user_openid", _taobaoUserOpenid)
	return nil
}

// GetTaobaoUserOpenid TaobaoUserOpenid Getter
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetTaobaoUserOpenid() string {
	return r._taobaoUserOpenid
}

// SetExtInfo is ExtInfo Setter
// 扩展数据
func (r *AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r AlibabaailabstmallgenieauthdeviceqrcodeactivateAPIRequest) GetExtInfo() string {
	return r._extInfo
}
