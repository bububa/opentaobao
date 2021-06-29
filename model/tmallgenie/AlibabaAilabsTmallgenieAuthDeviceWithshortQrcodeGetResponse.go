package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据安全简码查询二维码详细信息 APIResponse
alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get

根据安全简码查询二维码详细信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetResponse
}

type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withshort_qrcode_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 二维码数据
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
