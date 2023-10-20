package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse 根据userId(支持单个或批量)获取用户实时位置信息 API返回值
// alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids
//
// 根据userId(支持单个或批量)获取用户实时位置信息
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：getActualUserLocationInfoByIds
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponseModel).Reset()
}

// AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponseModel is 根据userId(支持单个或批量)获取用户实时位置信息 成功返回结果
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getactualuserlocationinfobyids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse)
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse
func GetAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse() *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse {
	return poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse.Get().(*AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse 将 AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse(v *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIResponse.Put(v)
}
