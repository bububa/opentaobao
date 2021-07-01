package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据三方ID查询设备注册激活信息 API返回值 
alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get

根据三方ID查询设备注册激活信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel
}

// 根据三方ID查询设备注册激活信息 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withdeviceid_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
