package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionPublishAPIResponse 发布应用升级 API返回值
// yunos.osupdate.appversion.publish
//
// 发布应用升级任务
type YunosOsupdateAppversionPublishAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionPublishAPIResponseModel
}

// YunosOsupdateAppversionPublishAPIResponseModel is 发布应用升级 成功返回结果
type YunosOsupdateAppversionPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosOsupdateAppversionPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}
