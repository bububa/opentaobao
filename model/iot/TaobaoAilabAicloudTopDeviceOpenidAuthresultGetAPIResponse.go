package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponse 获取openId设备授权码验证结果 API返回值
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
type TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponseModel
}

// TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponseModel is 获取openId设备授权码验证结果 成功返回结果
type TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_openid_authresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
