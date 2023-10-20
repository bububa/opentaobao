package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceApksAPIResponse 获取停开服apk列表 API返回值
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
type YunosTvpubadminDeviceApksAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceApksAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceApksAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceApksAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceApksAPIResponseModel is 获取停开服apk列表 成功返回结果
type YunosTvpubadminDeviceApksAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_apks_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	ApkList []DicControlApkDo `json:"apk_list,omitempty" xml:"apk_list>dic_control_apk_do,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceApksAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApkList = m.ApkList[:0]
}

var poolYunosTvpubadminDeviceApksAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceApksAPIResponse)
	},
}

// GetYunosTvpubadminDeviceApksAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceApksAPIResponse
func GetYunosTvpubadminDeviceApksAPIResponse() *YunosTvpubadminDeviceApksAPIResponse {
	return poolYunosTvpubadminDeviceApksAPIResponse.Get().(*YunosTvpubadminDeviceApksAPIResponse)
}

// ReleaseYunosTvpubadminDeviceApksAPIResponse 将 YunosTvpubadminDeviceApksAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceApksAPIResponse(v *YunosTvpubadminDeviceApksAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceApksAPIResponse.Put(v)
}
