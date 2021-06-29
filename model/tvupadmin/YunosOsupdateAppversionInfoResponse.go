package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 API返回值 
yunos.osupdate.appversion.info

获取应用升级详情
*/
type YunosOsupdateAppversionInfoAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionInfoResponse
}

// 获取应用升级详情 成功返回结果
type YunosOsupdateAppversionInfoResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_info_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *TvAppVersion `json:"data,omitempty" xml:"data,omitempty"`
}
