package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码激活设备 APIRequest
alibaba.ailabs.tmallgenie.auth.device.qrcode.activate

三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。
*/
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest struct {
    model.Params

    // OAUTH authcode码
    code   string 

    // 产品终端ID
    clientId   string 

    // 淘宝授权ID
    taobaoUserOpenid   string 

    // 扩展数据
    extInfo   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest() *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.qrcode.activate"
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) SetTaobaoUserOpenid(taobaoUserOpenid string) error {
    r.taobaoUserOpenid = taobaoUserOpenid
    r.Set("taobao_user_openid", taobaoUserOpenid)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetTaobaoUserOpenid() string {
    return r.taobaoUserOpenid
}

func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) SetExtInfo(extInfo string) error {
    r.extInfo = extInfo
    r.Set("ext_info", extInfo)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateRequest) GetExtInfo() string {
    return r.extInfo
}

