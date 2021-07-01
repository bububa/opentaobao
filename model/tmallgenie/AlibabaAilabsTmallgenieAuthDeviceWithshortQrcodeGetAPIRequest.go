package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据安全简码查询二维码详细信息 API请求
alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get

根据安全简码查询二维码详细信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest struct {
    model.Params
    // 产品ID
    _clientId   string
    // 授权码
    _authCode   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 产品ID
func (r *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) GetClientId() string {
    return r._clientId
}
// AuthCode Setter
// 授权码
func (r *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) SetAuthCode(_authCode string) error {
    r._authCode = _authCode
    r.Set("auth_code", _authCode)
    return nil
}

// AuthCode Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest) GetAuthCode() string {
    return r._authCode
}
