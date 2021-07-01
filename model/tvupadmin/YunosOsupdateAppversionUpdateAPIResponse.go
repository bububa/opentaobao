package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionUpdateAPIResponse
应用升级任务更新 API返回值
yunos.osupdate.appversion.update

应用升级任务更新 */
type YunosOsupdateAppversionUpdateAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionUpdateAPIResponseModel
}

// YunosOsupdateAppversionUpdateAPIResponseModel is 应用升级任务更新 成功返回结果
type YunosOsupdateAppversionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
