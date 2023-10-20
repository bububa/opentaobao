package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest 扫码激活设备 API请求
// alibaba.ailabs.tmallgenie.auth.device.qrcode.activate
//
// 三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) Reset() {
	r._code = ""
	r._clientId = ""
	r._taobaoUserOpenid = ""
	r._extInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.qrcode.activate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// OAUTH authcode码
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetCode() string {
	return r._code
}

// SetClientId is ClientId Setter
// 产品终端ID
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetClientId() string {
	return r._clientId
}

// SetTaobaoUserOpenid is TaobaoUserOpenid Setter
// 淘宝授权ID
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetTaobaoUserOpenid(_taobaoUserOpenid string) error {
	r._taobaoUserOpenid = _taobaoUserOpenid
	r.Set("taobao_user_openid", _taobaoUserOpenid)
	return nil
}

// GetTaobaoUserOpenid TaobaoUserOpenid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetTaobaoUserOpenid() string {
	return r._taobaoUserOpenid
}

// SetExtInfo is ExtInfo Setter
// 扩展数据
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) GetExtInfo() string {
	return r._extInfo
}

var poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest.Put(v)
}
