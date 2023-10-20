package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponse 获取openid设备通用授权码 API返回值
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
type TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponseModel
}

// TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponseModel is 获取openid设备通用授权码 成功返回结果
type TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_openid_authcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
