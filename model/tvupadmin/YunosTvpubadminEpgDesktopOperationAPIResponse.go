package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminEpgDesktopOperationAPIResponse 影视桌面运营通用接口 API返回值
// yunos.tvpubadmin.epg.desktop.operation
//
// 影视桌面运营通用接口
type YunosTvpubadminEpgDesktopOperationAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminEpgDesktopOperationAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminEpgDesktopOperationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminEpgDesktopOperationAPIResponseModel).Reset()
}

// YunosTvpubadminEpgDesktopOperationAPIResponseModel is 影视桌面运营通用接口 成功返回结果
type YunosTvpubadminEpgDesktopOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_epg_desktop_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 具体返回结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminEpgDesktopOperationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolYunosTvpubadminEpgDesktopOperationAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminEpgDesktopOperationAPIResponse)
	},
}

// GetYunosTvpubadminEpgDesktopOperationAPIResponse 从 sync.Pool 获取 YunosTvpubadminEpgDesktopOperationAPIResponse
func GetYunosTvpubadminEpgDesktopOperationAPIResponse() *YunosTvpubadminEpgDesktopOperationAPIResponse {
	return poolYunosTvpubadminEpgDesktopOperationAPIResponse.Get().(*YunosTvpubadminEpgDesktopOperationAPIResponse)
}

// ReleaseYunosTvpubadminEpgDesktopOperationAPIResponse 将 YunosTvpubadminEpgDesktopOperationAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminEpgDesktopOperationAPIResponse(v *YunosTvpubadminEpgDesktopOperationAPIResponse) {
	v.Reset()
	poolYunosTvpubadminEpgDesktopOperationAPIResponse.Put(v)
}
