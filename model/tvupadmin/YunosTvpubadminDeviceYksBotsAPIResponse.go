package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备列表 API返回值 
yunos.tvpubadmin.device.yks.bots

获取设备列表
*/
type YunosTvpubadminDeviceYksBotsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksBotsAPIResponseModel
}

// 获取设备列表 成功返回结果
type YunosTvpubadminDeviceYksBotsAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_bots_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
