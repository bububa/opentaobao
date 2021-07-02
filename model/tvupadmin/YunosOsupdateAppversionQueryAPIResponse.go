package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionQueryAPIResponse 分页获取桌面升级任务 API返回值
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
type YunosOsupdateAppversionQueryAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionQueryAPIResponseModel
}

// YunosOsupdateAppversionQueryAPIResponseModel is 分页获取桌面升级任务 成功返回结果
type YunosOsupdateAppversionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultList
	ResultList []TvAppVersion `json:"result_list,omitempty" xml:"result_list>tv_app_version,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
