package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceAppupgradedetailAPIResponse 获取应用升级详情 API返回值
// yunos.tvpubadmin.device.appupgradedetail
//
// 获取应用升级详情
type YunosTvpubadminDeviceAppupgradedetailAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceAppupgradedetailAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceAppupgradedetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceAppupgradedetailAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceAppupgradedetailAPIResponseModel is 获取应用升级详情 成功返回结果
type YunosTvpubadminDeviceAppupgradedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_appupgradedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应用列表
	Object *AppVersionAuditDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceAppupgradedetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDeviceAppupgradedetailAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceAppupgradedetailAPIResponse)
	},
}

// GetYunosTvpubadminDeviceAppupgradedetailAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceAppupgradedetailAPIResponse
func GetYunosTvpubadminDeviceAppupgradedetailAPIResponse() *YunosTvpubadminDeviceAppupgradedetailAPIResponse {
	return poolYunosTvpubadminDeviceAppupgradedetailAPIResponse.Get().(*YunosTvpubadminDeviceAppupgradedetailAPIResponse)
}

// ReleaseYunosTvpubadminDeviceAppupgradedetailAPIResponse 将 YunosTvpubadminDeviceAppupgradedetailAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceAppupgradedetailAPIResponse(v *YunosTvpubadminDeviceAppupgradedetailAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceAppupgradedetailAPIResponse.Put(v)
}
