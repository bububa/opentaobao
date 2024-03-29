package mirage

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuMirageQueryPermissionAPIResponse 优酷播控查询是否可播API API返回值
// youku.mirage.query.permission
//
// 根据节目ID或者VID查询视频或者节目是否可以播放
type YoukuMirageQueryPermissionAPIResponse struct {
	model.CommonResponse
	YoukuMirageQueryPermissionAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuMirageQueryPermissionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuMirageQueryPermissionAPIResponseModel).Reset()
}

// YoukuMirageQueryPermissionAPIResponseModel is 优酷播控查询是否可播API 成功返回结果
type YoukuMirageQueryPermissionAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_mirage_query_permission_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Resp *PermissionResponseDto `json:"resp,omitempty" xml:"resp,omitempty"`
}

// Reset 清空结构体
func (m *YoukuMirageQueryPermissionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resp = nil
}

var poolYoukuMirageQueryPermissionAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuMirageQueryPermissionAPIResponse)
	},
}

// GetYoukuMirageQueryPermissionAPIResponse 从 sync.Pool 获取 YoukuMirageQueryPermissionAPIResponse
func GetYoukuMirageQueryPermissionAPIResponse() *YoukuMirageQueryPermissionAPIResponse {
	return poolYoukuMirageQueryPermissionAPIResponse.Get().(*YoukuMirageQueryPermissionAPIResponse)
}

// ReleaseYoukuMirageQueryPermissionAPIResponse 将 YoukuMirageQueryPermissionAPIResponse 保存到 sync.Pool
func ReleaseYoukuMirageQueryPermissionAPIResponse(v *YoukuMirageQueryPermissionAPIResponse) {
	v.Reset()
	poolYoukuMirageQueryPermissionAPIResponse.Put(v)
}
