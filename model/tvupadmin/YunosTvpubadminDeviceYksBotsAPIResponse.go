package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksBotsAPIResponse 获取设备列表 API返回值
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
type YunosTvpubadminDeviceYksBotsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksBotsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksBotsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceYksBotsAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceYksBotsAPIResponseModel is 获取设备列表 成功返回结果
type YunosTvpubadminDeviceYksBotsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_bots_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksBotsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDeviceYksBotsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceYksBotsAPIResponse)
	},
}

// GetYunosTvpubadminDeviceYksBotsAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceYksBotsAPIResponse
func GetYunosTvpubadminDeviceYksBotsAPIResponse() *YunosTvpubadminDeviceYksBotsAPIResponse {
	return poolYunosTvpubadminDeviceYksBotsAPIResponse.Get().(*YunosTvpubadminDeviceYksBotsAPIResponse)
}

// ReleaseYunosTvpubadminDeviceYksBotsAPIResponse 将 YunosTvpubadminDeviceYksBotsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceYksBotsAPIResponse(v *YunosTvpubadminDeviceYksBotsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceYksBotsAPIResponse.Put(v)
}
