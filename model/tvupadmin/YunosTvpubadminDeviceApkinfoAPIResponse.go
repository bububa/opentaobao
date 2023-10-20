package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceApkinfoAPIResponse 获取停开服apk信息 API返回值
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
type YunosTvpubadminDeviceApkinfoAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceApkinfoAPIResponseModel
}

// YunosTvpubadminDeviceApkinfoAPIResponseModel is 获取停开服apk信息 成功返回结果
type YunosTvpubadminDeviceApkinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_apkinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *DicControlApkDo `json:"object,omitempty" xml:"object,omitempty"`
}
