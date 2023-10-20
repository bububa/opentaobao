package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceOsupgradequeryAPIResponse 系统升级查询 API返回值
// yunos.tvpubadmin.device.osupgradequery
//
// 系统升级查询
type YunosTvpubadminDeviceOsupgradequeryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceOsupgradequeryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceOsupgradequeryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceOsupgradequeryAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceOsupgradequeryAPIResponseModel is 系统升级查询 成功返回结果
type YunosTvpubadminDeviceOsupgradequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_osupgradequery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 信息结构
	ObjectList *PaginationDo `json:"object_list,omitempty" xml:"object_list,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceOsupgradequeryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ObjectList = nil
}

var poolYunosTvpubadminDeviceOsupgradequeryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceOsupgradequeryAPIResponse)
	},
}

// GetYunosTvpubadminDeviceOsupgradequeryAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceOsupgradequeryAPIResponse
func GetYunosTvpubadminDeviceOsupgradequeryAPIResponse() *YunosTvpubadminDeviceOsupgradequeryAPIResponse {
	return poolYunosTvpubadminDeviceOsupgradequeryAPIResponse.Get().(*YunosTvpubadminDeviceOsupgradequeryAPIResponse)
}

// ReleaseYunosTvpubadminDeviceOsupgradequeryAPIResponse 将 YunosTvpubadminDeviceOsupgradequeryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceOsupgradequeryAPIResponse(v *YunosTvpubadminDeviceOsupgradequeryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceOsupgradequeryAPIResponse.Put(v)
}
