package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceStatsAPIResponse 获取设备统计数据 API返回值
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
type YunosTvpubadminDeviceStatsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceStatsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceStatsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceStatsAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceStatsAPIResponseModel is 获取设备统计数据 成功返回结果
type YunosTvpubadminDeviceStatsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_stats_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *StatsDeviceInfoDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceStatsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDeviceStatsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceStatsAPIResponse)
	},
}

// GetYunosTvpubadminDeviceStatsAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceStatsAPIResponse
func GetYunosTvpubadminDeviceStatsAPIResponse() *YunosTvpubadminDeviceStatsAPIResponse {
	return poolYunosTvpubadminDeviceStatsAPIResponse.Get().(*YunosTvpubadminDeviceStatsAPIResponse)
}

// ReleaseYunosTvpubadminDeviceStatsAPIResponse 将 YunosTvpubadminDeviceStatsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceStatsAPIResponse(v *YunosTvpubadminDeviceStatsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceStatsAPIResponse.Put(v)
}
