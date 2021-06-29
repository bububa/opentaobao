package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据mac查询设备的安全二维码 APIRequest
alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get

根据mac查询二维码详细信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest struct {
    model.Params

    // 产品ID
    clientId   string 

    // 设备mac地址
    mac   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get"
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest) GetMac() string {
    return r.mac
}

