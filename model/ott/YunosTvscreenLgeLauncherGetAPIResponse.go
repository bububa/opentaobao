package ott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvscreenlgelaunchergetAPIResponse LG用桌面信息获取 API返回值
// yunos.tvscreen.lge.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务,根据LG的需求定制输出格式
type YunostvscreenlgelaunchergetAPIResponse struct {
	model.CommonResponse
	YunostvscreenlgelaunchergetAPIResponseModel
}

// YunostvscreenlgelaunchergetAPIResponseModel is LG用桌面信息获取 成功返回结果
type YunostvscreenlgelaunchergetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvscreen_lge_launcher_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunostvscreenlgelaunchergetResult `json:"result,omitempty" xml:"result,omitempty"`
}
