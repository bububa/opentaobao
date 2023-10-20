package ott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvscreenLgeLauncherGetAPIResponse LG用桌面信息获取 API返回值
// yunos.tvscreen.lge.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务,根据LG的需求定制输出格式
type YunosTvscreenLgeLauncherGetAPIResponse struct {
	model.CommonResponse
	YunosTvscreenLgeLauncherGetAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvscreenLgeLauncherGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvscreenLgeLauncherGetAPIResponseModel).Reset()
}

// YunosTvscreenLgeLauncherGetAPIResponseModel is LG用桌面信息获取 成功返回结果
type YunosTvscreenLgeLauncherGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvscreen_lge_launcher_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunosTvscreenLgeLauncherGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvscreenLgeLauncherGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvscreenLgeLauncherGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvscreenLgeLauncherGetAPIResponse)
	},
}

// GetYunosTvscreenLgeLauncherGetAPIResponse 从 sync.Pool 获取 YunosTvscreenLgeLauncherGetAPIResponse
func GetYunosTvscreenLgeLauncherGetAPIResponse() *YunosTvscreenLgeLauncherGetAPIResponse {
	return poolYunosTvscreenLgeLauncherGetAPIResponse.Get().(*YunosTvscreenLgeLauncherGetAPIResponse)
}

// ReleaseYunosTvscreenLgeLauncherGetAPIResponse 将 YunosTvscreenLgeLauncherGetAPIResponse 保存到 sync.Pool
func ReleaseYunosTvscreenLgeLauncherGetAPIResponse(v *YunosTvscreenLgeLauncherGetAPIResponse) {
	v.Reset()
	poolYunosTvscreenLgeLauncherGetAPIResponse.Put(v)
}
