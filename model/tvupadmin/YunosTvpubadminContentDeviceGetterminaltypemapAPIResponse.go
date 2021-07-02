package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse 获取终端类型表 API返回值
// yunos.tvpubadmin.content.device.getterminaltypemap
//
// 获取终端类型表
type YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentDeviceGetterminaltypemapAPIResponseModel
}

// YunosTvpubadminContentDeviceGetterminaltypemapAPIResponseModel is 获取终端类型表 成功返回结果
type YunosTvpubadminContentDeviceGetterminaltypemapAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_device_getterminaltypemap_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// map
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
