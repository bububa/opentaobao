package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest
扫码激活设备 API请求
alibaba.ailabs.tmallgenie.auth.device.qrcode.activate

三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。 */
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest struct {
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

// NewAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest 初始化AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest() *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.qrcode.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// OAUTH authcode码
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetCode() string {
	return r._code
}

// Set is ClientId Setter
// 产品终端ID
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// Get ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetClientId() string {
	return r._clientId
}

// Set is TaobaoUserOpenid Setter
// 淘宝授权ID
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetTaobaoUserOpenid(_taobaoUserOpenid string) error {
	r._taobaoUserOpenid = _taobaoUserOpenid
	r.Set("taobao_user_openid", _taobaoUserOpenid)
	return nil
}

// Get TaobaoUserOpenid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetTaobaoUserOpenid() string {
	return r._taobaoUserOpenid
}

// Set is ExtInfo Setter
// 扩展数据
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// Get ExtInfo Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetExtInfo() string {
	return r._extInfo
}
