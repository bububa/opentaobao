package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse 分时间段获取用户历史位置信息 API返回值
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponseModel).Reset()
}

// AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponseModel is 分时间段获取用户历史位置信息 成功返回结果
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getuserlocationinfologs_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse)
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse
func GetAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse() *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse {
	return poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse.Get().(*AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse 将 AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse(v *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIResponse.Put(v)
}
