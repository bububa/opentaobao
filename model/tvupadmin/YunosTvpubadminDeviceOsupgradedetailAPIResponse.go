package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceOsupgradedetailAPIResponse 获取系统升级详情 API返回值
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
type YunosTvpubadminDeviceOsupgradedetailAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceOsupgradedetailAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceOsupgradedetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceOsupgradedetailAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceOsupgradedetailAPIResponseModel is 获取系统升级详情 成功返回结果
type YunosTvpubadminDeviceOsupgradedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_osupgradedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 具体的数据对象
	Object *OsVersionAuditDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceOsupgradedetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDeviceOsupgradedetailAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceOsupgradedetailAPIResponse)
	},
}

// GetYunosTvpubadminDeviceOsupgradedetailAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceOsupgradedetailAPIResponse
func GetYunosTvpubadminDeviceOsupgradedetailAPIResponse() *YunosTvpubadminDeviceOsupgradedetailAPIResponse {
	return poolYunosTvpubadminDeviceOsupgradedetailAPIResponse.Get().(*YunosTvpubadminDeviceOsupgradedetailAPIResponse)
}

// ReleaseYunosTvpubadminDeviceOsupgradedetailAPIResponse 将 YunosTvpubadminDeviceOsupgradedetailAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceOsupgradedetailAPIResponse(v *YunosTvpubadminDeviceOsupgradedetailAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceOsupgradedetailAPIResponse.Put(v)
}
