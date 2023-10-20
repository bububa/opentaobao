package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse 根据分组条件查询分组下的空间单元不包涵业务属性信息 API返回值
// alibaba.campus.adminmap.poiinfo.getlistbygroup
//
// 根据分组条件查询分组下的空间单元不包涵业务属性信息
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponseModel).Reset()
}

// AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponseModel is 根据分组条件查询分组下的空间单元不包涵业务属性信息 成功返回结果
type AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_poiinfo_getlistbygroup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse)
	},
}

// GetAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse 从 sync.Pool 获取 AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse
func GetAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse() *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse {
	return poolAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse.Get().(*AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse)
}

// ReleaseAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse 将 AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse(v *AlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse) {
	v.Reset()
	poolAlibabaCampusAdminmapPoiinfoGetlistbygroupAPIResponse.Put(v)
}
