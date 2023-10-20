package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateosfotapublishAPIResponse 系统升级发布 API返回值
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
type YunososupdateosfotapublishAPIResponse struct {
	model.CommonResponse
	YunososupdateosfotapublishAPIResponseModel
}

// YunososupdateosfotapublishAPIResponseModel is 系统升级发布 成功返回结果
type YunososupdateosfotapublishAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_osfota_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunososupdateosfotapublishResult `json:"result,omitempty" xml:"result,omitempty"`
}
