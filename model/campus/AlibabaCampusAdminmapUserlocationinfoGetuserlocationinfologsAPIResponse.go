package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponse 分时间段获取用户历史位置信息 API返回值
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
type AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponse struct {
	model.CommonResponse
	AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponseModel
}

// AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponseModel is 分时间段获取用户历史位置信息 成功返回结果
type AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getuserlocationinfologs_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
