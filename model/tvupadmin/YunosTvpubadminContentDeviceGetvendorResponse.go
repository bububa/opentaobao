package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备Vendor信息 API返回值 
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息
*/
type YunosTvpubadminContentDeviceGetvendorAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentDeviceGetvendorResponse
}

// 查询设备Vendor信息 成功返回结果
type YunosTvpubadminContentDeviceGetvendorResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_device_getvendor_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // list
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
